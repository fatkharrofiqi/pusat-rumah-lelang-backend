package interfaces

import "io"

type IStorageGeneric interface {
	UploadFile(objectName string, filePath string) error
	GetFile(objectName string, writer io.Writer) error
	RemoveFile(objectName string) error
	ListFiles(bucketName string, prefix string) ([]string, error)
	GetFileURL(objectName string, expiry int64) (string, error)
}
