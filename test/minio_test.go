package test

import (
	"fmt"
	"path/filepath"
	"pusat-rumah-lelang-backend/config"
	"pusat-rumah-lelang-backend/constants"
	"pusat-rumah-lelang-backend/helpers"
	"sync"
	"testing"

	"github.com/k0kubun/pp/v3"
)

func setup() (minio *helpers.MinioStorage, err error) {
	minio, err = config.NewMinioStorage()
	if err != nil {
		pp.Println("Error creating minio storage", err.Error())
	}

	return
}

func upload(wg *sync.WaitGroup, minio *helpers.MinioStorage, key string, index int, photoPath string) error {
	wg.Add(1)
	defer wg.Done()
	namefile, err := minio.UploadFile(fmt.Sprintf("%s/%s/%d", "photo_house", key, index), filepath.Join(constants.RootDir, "/data/photo", photoPath))
	if err != nil {
		return err
	}
	pp.Println(namefile)
	return nil
}

func TestUploadMinio(t *testing.T) {
	config.LoadEnv()
	config.LoadConstant()
	minio, err := setup()
	if err != nil {
		pp.Fatal(err.Error())
	}

	t.Run("upload", func(t *testing.T) {
		photosPaths, err := helpers.RetrieveFiles(filepath.Join(constants.RootDir + "/data/photo"))
		if err != nil {
			pp.Fatal(err.Error())
		}
		var wg sync.WaitGroup
		for key, house := range photosPaths {
			for index, photoPath := range house {
				go upload(&wg, minio, key, index, photoPath)
			}
		}
		wg.Wait()
	})

	// t.Run("getFileUrl", func(t *testing.T) {
	// 	url, err := minio.GetFileURL("folder/file.jpeg", 70)
	// 	if err != nil {
	// 		pp.Println(err.Error())
	// 		panic(err.Error())
	// 	}

	// 	pp.Println(url)
	// })

	// t.Run("delete file", func(t *testing.T) {
	// 	err := minio.RemoveFile("folder/file.jpeg")
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// })
}
