package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
)

func HashKey(key string) []byte {
	hash := sha256.Sum256([]byte("maulanamuhammad"))
	return hash[:]
}

// Fungsi untuk mengenkripsi teks
func Encrypt(plaintext string) (string, error) {
	key := HashKey("maulanamuhammad")
	// Membuat blok cipher AES
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Membuat nonce untuk GCM
	nonce := make([]byte, 12) // Nonce berukuran 12 byte sesuai standar GCM
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// Membuat AES-GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Mengenkripsi plaintext
	ciphertext := aesGCM.Seal(nil, nonce, []byte(plaintext), nil)

	// Menggabungkan nonce dan ciphertext, lalu meng-encode ke base64
	result := append(nonce, ciphertext...)
	return base64.StdEncoding.EncodeToString(result), nil
}

func Decrypt(encryptedString string) (string, error) {
	keyString := HashKey("maulanamuhammad")
	// key, _ := hex.DecodeString(keyString)
	enc, _ := base64.StdEncoding.DecodeString(encryptedString)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher([]byte(keyString))
	if err != nil {
		return "", err
	}

	//Create a new GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	//Get the nonce size
	nonceSize := aesGCM.NonceSize()

	//Extract the nonce from the encrypted data
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	//Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s", plaintext), nil
}
