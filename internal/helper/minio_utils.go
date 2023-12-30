package helper

import (
	"context"
	"errors"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/minio/minio-go/v7"
)

type MinioStorage struct {
	Client     *minio.Client
	BucketName string
}

func (ms *MinioStorage) UploadFile(objectName string, filePath string) (string, error) {
	// Open the file.
	file, err := os.Open(filePath)
	if err != nil {
		return "", errors.New("Can't open file: " + filePath + ": " + err.Error())
	}
	defer file.Close()

	// Get file stats.
	fileInfo, err := file.Stat()
	if err != nil {
		return "", errors.New("Can't stat file: " + filePath)
	}

	// Custom object name - combine objectName with extension
	customObjectName := objectName + filepath.Ext(filePath)

	// Create an upload context with a cancellation signal.
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Upload the file using the custom object name.
	_, err = ms.Client.PutObject(ctx, ms.BucketName, customObjectName, file, fileInfo.Size(), minio.PutObjectOptions{
		ContentType: "image/jpeg",
		UserMetadata: map[string]string{
			"x-amz-acl": "public-read", // Set the ACL to make the object publicly readable
		},
		ContentDisposition: "inline",
	})
	if err != nil {
		return "", errors.New("Can't upload file with error: " + err.Error())
	}

	return customObjectName, nil
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
	err := ms.Client.RemoveObject(ctx, ms.BucketName, objectName, minio.RemoveObjectOptions{})
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
	reqParams := make(url.Values)
	reqParams.Set("response-content-type", "image/jpeg")
	reqParams.Set("response-content-disposition", "inline; filename="+objectName)
	presignedURL, err := ms.Client.PresignedGetObject(context.Background(), ms.BucketName, objectName, time.Duration(expiry)*time.Second, reqParams)
	if err != nil {
		return "", err
	}

	return presignedURL.String(), nil
}
