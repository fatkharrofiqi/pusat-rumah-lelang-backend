package config

import (
	"pusat-rumah-lelang-backend/internal/helper"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewMinioStorage(viper *viper.Viper) (*helper.MinioStorage, error) {
	endpoint := viper.GetString("MINIO_ENDPOINT")
	accessKey := viper.GetString("MINIO_ACCESS_KEY")
	secretKey := viper.GetString("MINIO_SECRET_KEY")
	ssl := viper.GetBool("MINIO_SSL")
	bucket := viper.GetString("MINIO_BUCKET")
	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: ssl,
	})
	if err != nil {
		return nil, err
	}

	return &helper.MinioStorage{
		Client:     minioClient,
		BucketName: bucket,
	}, nil
}

func NewMinio(viper *viper.Viper, log *logrus.Logger) *minio.Client {
	endpoint := viper.GetString("MINIO_ENDPOINT")
	accessKey := viper.GetString("MINIO_ACCESS_KEY")
	secretKey := viper.GetString("MINIO_SECRET_KEY")
	ssl := viper.GetBool("MINIO_SSL")
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: ssl,
	})
	if err != nil {
		log.Fatalf("Failed to create minio client: %v", err)
	}
	return minioClient
}
