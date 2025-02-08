package cloud

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/abhishekkushwahaa/secure-cloud-cli/internal/encryptor"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
)

func loadAWSConfig() (aws.Config, error) {
	_ = godotenv.Load()

	bucketName := os.Getenv("S3_BUCKET")
	accessKey := os.Getenv("ACCESS_KEY")
	secretKey := os.Getenv("SECRET_KEY")
	region := os.Getenv("REGION")

	if bucketName == "" || accessKey == "" || secretKey == "" || region == "" {
		return aws.Config{}, fmt.Errorf("missing AWS credentials in environment variables")
	}

	return config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
	)
}

func EnsureKeyGenerate() ([]byte, error) {
	key, err := encryptor.LoadKey()
	if err == nil {
		return key, nil
	}

	key, err = encryptor.GenerateKey()
	if err != nil {
		return nil, err
	}

	if err := encryptor.SaveKey(key); err != nil {
		return nil, err
	}
	return key, nil
}

func UploadToS3(filename string) error {
	// Check if the file exists before proceeding
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return fmt.Errorf("file %s does not exist", filename)
	}

	cfg, err := loadAWSConfig()
	if err != nil {
		return err
	}

	key, err := EnsureKeyGenerate()
	if err != nil {
		return err
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	encryptedData, err := encryptor.Encrypt(data, key)
	if err != nil {
		return err
	}

	// Extract only the filename (without directory)
	filePath := filepath.Base(filename)

	client := s3.NewFromConfig(cfg)
	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET")),
		Key:    aws.String(filePath),
		Body:   bytes.NewReader([]byte(encryptedData)),
	})

	return err
}

func DownloadFromS3(filename string) error {
	cfg, err := loadAWSConfig()
	if err != nil {
		return err
	}

	key, err := encryptor.LoadKey()
	if err != nil {
		return err
	}

	client := s3.NewFromConfig(cfg)
	output, err := client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET")),
		Key:    aws.String(filename),
	})
	if err != nil {
		return err
	}
	defer output.Body.Close()

	// Read the encrypted file from S3
	buffer := new(bytes.Buffer)
	_, err = buffer.ReadFrom(output.Body)
	if err != nil {
		return err
	}

	decryptedData, err := encryptor.Decrypt(buffer.String(), key)
	if err != nil {
		return err
	}

	savePath := "data/downloads/" + filename
	err = os.WriteFile(savePath, decryptedData, 0600)
	if err != nil {
		return err
	}
	return nil
}
