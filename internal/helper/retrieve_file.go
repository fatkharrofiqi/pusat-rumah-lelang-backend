package helper

import (
	"fmt"
	"image"
	_ "image/jpeg" // JPEG format support
	_ "image/png"  // PNG format support
	"log"
	"os"
	"path/filepath"
)

func RetrieveFiles(root string) (map[string][]string, error) {
	filePaths := make(map[string][]string)

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error accessing path %q: %v\n", path, err)
			return err
		}
		if !info.IsDir() {
			dir := filepath.Dir(path)
			relPath, err := filepath.Rel(root, path)
			if err != nil {
				return err
			}
			baseDir := filepath.Base(dir) // Get base name of directory
			filePaths[baseDir] = append(filePaths[baseDir], relPath)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return filePaths, nil
}

func GetSize(filePath string) (width int, height int) {
	// Open the image file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()
	// Decode the image file
	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Get the dimensions (width and height) of the image
	bounds := img.Bounds()
	width = bounds.Dx()  // Width of the image
	height = bounds.Dy() // Height of the image
	return
}
