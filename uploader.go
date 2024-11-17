package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
)

// Uploader interface
type Uploader interface {
	UploadFile(ctx context.Context, file multipart.File, fileName string) (string, error)
}

// S3Uploader implements Uploader interface
type S3Uploader struct {
	s3Client *s3.S3
	bucket   string
}

// NewS3Uploader creates a new S3Uploader with access keys
func NewS3Uploader(region, bucket, accessKeyID, secretAccessKey string) *S3Uploader {
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKeyID, secretAccessKey, ""),
	}))
	return &S3Uploader{
		s3Client: s3.New(sess),
		bucket:   bucket,
	}
}

// UploadFile uploads a file to S3 with a UUID as its name
func (u *S3Uploader) UploadFile(ctx context.Context, file multipart.File, originalFileName string) (string, error) {
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	var buf bytes.Buffer
	_, err := io.Copy(&buf, file)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %v", err)
	}

	fileExtension := getFileExtension(originalFileName)
	newFileName := uuid.New().String() + fileExtension

	_, err = u.s3Client.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(u.bucket),
		Key:         aws.String(newFileName),
		Body:        bytes.NewReader(buf.Bytes()),
		ContentType: aws.String("application/octet-stream"),
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file: %v", err)
	}

	url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", u.bucket, *u.s3Client.Config.Region, newFileName)
	return url, nil
}

func getFileExtension(fileName string) string {
	ext := ""
	for i := len(fileName) - 1; i >= 0 && fileName[i] != '.' && fileName[i] != '/'; i-- {
		ext = string(fileName[i]) + ext
	}
	if ext != "" {
		return "." + ext
	}
	return ""
}
