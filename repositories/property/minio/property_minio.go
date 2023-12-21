package minio

import (
	"context"
	"os"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioStorage struct {
	Client *minio.Client
}

func NewMinioStorage(endpoint, accessKey, secretKey string, useSSL bool) (*MinioStorage, error) {
	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, err
	}

	return &MinioStorage{
		Client: minioClient,
	}, nil
}

func (ms *MinioStorage) UploadFile(objectName string, filePath string) error {
	// Open the file.
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Get file stats.
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	// Create an upload context with a cancellation signal.
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Upload the file.
	_, err = ms.Client.PutObject(ctx, "prl", objectName, file, fileInfo.Size(), minio.PutObjectOptions{})
	if err != nil {
		return err
	}

	return nil
}

// func (ms *MinioStorage) GetFile(objectName string, writer io.Writer) error {
// 	// Create a download context with a cancellation signal.
// 	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
// 	defer cancel()

// 	// Download the file.
// 	err := ms.Client.FGetObject(ctx, "your-bucket-name", objectName, writer, minio.GetObjectOptions{})
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func (ms *MinioStorage) RemoveFile(objectName string) error {
	// Create a delete context with a cancellation signal.
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Remove the file.
	err := ms.Client.RemoveObject(ctx, "your-bucket-name", objectName, minio.RemoveObjectOptions{})
	if err != nil {
		return err
	}

	return nil
}

// func (ms *MinioStorage) ListFiles(bucketName string, prefix string) ([]string, error) {
// 	// Create a context.
// 	ctx := context.Background()

// 	// List objects in the bucket with the specified prefix.
// 	doneCh := make(chan struct{})
// 	defer close(doneCh)

// 	var objects []string
// 	for object := range ms.Client.ListObjects(ctx, bucketName, minio.ListObjectsOptions{Prefix: prefix, Recursive: true}, doneCh) {
// 		if object.Err != nil {
// 			return nil, object.Err
// 		}
// 		objects = append(objects, object.Key)
// 	}

// 	return objects, nil
// }

func (ms *MinioStorage) GetFileURL(objectName string, expiry int64) (string, error) {
	// Generate a presigned URL for the object with expiry time.
	presignedURL, err := ms.Client.PresignedGetObject(context.Background(), "your-bucket-name", objectName, time.Duration(expiry)*time.Second, nil)
	if err != nil {
		return "", err
	}

	return presignedURL.String(), nil
}
