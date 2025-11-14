package fileupload

import (
	"context"
	"echo-boilerplate/config"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	s3types "github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/labstack/echo/v4"
)

func GetFileHandler(c echo.Context) error {
	list, _ := config.S3Client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(os.Getenv("AWS_BUCKET")),
	})

	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
		"message": "Get File Success",
		"data":    list,
	})
}

func UploadFileHandler(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	src, err := file.Open()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	defer src.Close()

	uploader := s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("AWS_BUCKET")),
		Key:    aws.String(file.Filename),
		Body:   src,
	}

	_, err = config.S3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:            uploader.Bucket,
		Key:               uploader.Key,
		Body:              uploader.Body,
		ACL:               "public-read",
		ChecksumAlgorithm: "",
		ChecksumSHA256:    aws.String(""),
	})
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("upload failed: %v", err))
	}

	fmt.Println(uploader)
	return c.JSON(http.StatusOK, map[string]string{
		"message": "upload success",
	})
}

func UploadPartInitHandler(c echo.Context) error {
	key := c.FormValue("key")
	uploadID, err := config.S3Client.CreateMultipartUpload(context.TODO(), &s3.CreateMultipartUploadInput{
		Bucket: aws.String(os.Getenv("AWS_BUCKET")),
		Key:    aws.String(key),
		ACL:    s3types.ObjectCannedACL("public-read"),
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// Save into database.

	return c.JSON(http.StatusOK, echo.Map{
		"upload_id": *uploadID.UploadId,
		"key":       key,
	})
}

// Step 2: Upload per part
func UploadPartHandler(c echo.Context) error {
	uploadID := c.FormValue("upload_id")
	partNumberStr := c.FormValue("part_number")
	key := c.FormValue("key")

	partNumber, _ := strconv.Atoi(partNumberStr)
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid file"})
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	defer src.Close()

	input := &s3.UploadPartInput{
		Bucket:         aws.String(os.Getenv("AWS_BUCKET")),
		Key:            aws.String(key),
		PartNumber:     aws.Int32(int32(partNumber)),
		UploadId:       aws.String(uploadID),
		Body:           src,
		ChecksumSHA256: aws.String(""),
	}

	result, err := config.S3Client.UploadPart(context.TODO(), input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "upload part success",
		"result":  result,
	})
}

func UploadPartCompleteHandler(c echo.Context) error {
	var req struct {
		UploadID string `json:"upload_id"`
		Key      string `json:"key"`
		Parts    []struct {
			PartNumber int32  `json:"PartNumber"`
			ETag       string `json:"ETag"`
		} `json:"parts"`
	}

	// Parse JSON body
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid json"})
	}

	// Build CompletedPart list
	var completedParts []s3types.CompletedPart
	for _, p := range req.Parts {
		completedParts = append(completedParts, s3types.CompletedPart{
			ETag:       aws.String(p.ETag),
			PartNumber: aws.Int32(p.PartNumber),
		})
	}

	// return c.JSON(http.StatusOK, completedParts)
	// Prepare complete request
	input := &s3.CompleteMultipartUploadInput{
		Bucket:   aws.String(os.Getenv("AWS_BUCKET")),
		Key:      aws.String(req.Key),
		UploadId: aws.String(req.UploadID),
		MultipartUpload: &s3types.CompletedMultipartUpload{
			Parts: completedParts,
		},

		ChecksumSHA256: aws.String(""), // disable sha256 mismatch
	}

	// Call S3
	result, err := config.S3Client.CompleteMultipartUpload(context.TODO(), input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "upload part complete success",
		"result":  result,
	})
}
