package config

import (
	"pusat-rumah-lelang-backend/constants"
	"pusat-rumah-lelang-backend/helpers"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func NewMinioStorage() (*helpers.MinioStorage, error) {
	// Initialize minio client object.
	minioClient, err := minio.New(constants.MinioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(constants.MinioAccessKeyId, constants.MinioSecretAccessKey, ""),
		Secure: constants.MinioUseSSL,
	})
	if err != nil {
		return nil, err
	}

	return &helpers.MinioStorage{
		Client:     minioClient,
		BucketName: constants.BucketName,
	}, nil
}
