package test

import (
	"fmt"
	"pusat-rumah-lelang-backend/internal/helper"
	"testing"
)

func TestRetriveFile(t *testing.T) {
	photoPath := "../data/photo"
	var filePaths map[string][]string
	var err error
	t.Run("RetrieveFile", func(t *testing.T) {
		filePaths, err = helper.RetrieveFiles(photoPath)
		if err != nil {
			panic(err.Error())
		}
	})

	t.Run("Loop through", func(t *testing.T) {
		for folder, files := range filePaths {
			fmt.Println(folder)
			for _, file := range files {
				fmt.Println(" -", file)
			}
		}
	})
}
