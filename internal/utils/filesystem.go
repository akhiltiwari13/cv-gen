package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// LoadFile loads a file and returns its contents as a string
func LoadFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("error opening file %s: %v", path, err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("error reading file %s: %v", path, err)
	}

	return string(content), nil
}

// EnsureDirExists ensures that a directory exists, creating it if necessary
func EnsureDirExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(path, 0755); err != nil {
			return fmt.Errorf("error creating directory %s: %v", path, err)
		}
	}
	return nil
}

// CopyFile copies a file from source to destination
func CopyFile(src, dst string) error {
	// Ensure destination directory exists
	dstDir := filepath.Dir(dst)
	if err := EnsureDirExists(dstDir); err != nil {
		return err
	}

	// Open source file
	srcFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("error opening source file %s: %v", src, err)
	}
	defer srcFile.Close()

	// Create destination file
	dstFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("error creating destination file %s: %v", dst, err)
	}
	defer dstFile.Close()

	// Copy contents
	if _, err = io.Copy(dstFile, srcFile); err != nil {
		return fmt.Errorf("error copying file contents: %v", err)
	}

	return nil
}

// FileExists checks if a file exists
func FileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// DirExists checks if a directory exists
func DirExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}
