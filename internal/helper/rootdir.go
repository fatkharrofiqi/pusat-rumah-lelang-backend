package helper

import (
	"errors"
	"path/filepath"
	"runtime"
)

func GetRootDir() (string, error) {
	// Get the path to the current file
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return "", errors.New("unable to get current file path")
	}

	// Deduce the root directory by traversing up the directory tree
	dir := filepath.Dir(filename)
	// Navigate up one directory level
	rootdir := filepath.Dir(dir)

	return rootdir, nil
}
