package helper

import (
	"fmt"
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
