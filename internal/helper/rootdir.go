package helper

import (
	"log"
	"path/filepath"
	"runtime"
)

func GetRootDir() string {
	// Get the path to the current file
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatalln("failed to get root directory")
	}

	// Deduce the root directory by traversing up the directory tree
	dir := filepath.Dir(filename)
	// Navigate up one directory level
	dir = filepath.Dir(dir)
	// Navigate up one directory level
	rootdir := filepath.Dir(dir)

	return rootdir
}
