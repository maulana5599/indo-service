package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var S3Client *s3.Client

type StorageConfig struct {
	AccessKey string
	SecretKey string
	Host      string
	Region    string
	Bucket    string
}

func LoadStorageConfig() StorageConfig {
	var config = StorageConfig{
		AccessKey: os.Getenv("AWS_ACCESS_KEY_ID"),
		SecretKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
		Host:      os.Getenv("AWS_ENDPOINT"),
		Region:    os.Getenv("AWS_REGION"),
		Bucket:    os.Getenv("AWS_BUCKET"),
	}

	return config
}

func StorageConnection() {
	var envConfig = LoadStorageConfig()

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			envConfig.AccessKey, // ganti
			envConfig.SecretKey, // ganti
			"",
		)),
		config.WithRegion(envConfig.Region), // sesuaikan region
		config.WithEndpointResolver(aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
			return aws.Endpoint{
				URL:               envConfig.Host, // ganti URL endpoint custom
				SigningRegion:     envConfig.Region,
				HostnameImmutable: true,
			}, nil
		})),
	)
	if err != nil {
		log.Fatalf("failed to load AWS config: %v", err)
	}

	fmt.Println("Connect storage successfully")
	S3Client = s3.NewFromConfig(cfg, func(options *s3.Options) {
		options.UsePathStyle = true
	})
}
