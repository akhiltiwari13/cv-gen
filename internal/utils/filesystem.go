// internal/utils/filesystem.go
package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/akhiltiwari13/cv-gen/internal/logging"
)

// LoadFile loads a file and returns its contents as a string
func LoadFile(path string) (string, error) {
	logger := logging.GetLogger()
	logger.Debug().Str("path", path).Msg("Loading file")

	file, err := os.Open(path)
	if err != nil {
		logger.Error().Err(err).Str("path", path).Msg("Error opening file")
		return "", fmt.Errorf("error opening file %s: %v", path, err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		logger.Error().Err(err).Str("path", path).Msg("Error reading file contents")
		return "", fmt.Errorf("error reading file %s: %v", path, err)
	}

	logger.Debug().Str("path", path).Int("content_length", len(content)).Msg("File loaded successfully")
	return string(content), nil
}

// EnsureDirExists ensures that a directory exists, creating it if necessary
func EnsureDirExists(path string) error {
	logger := logging.GetLogger()
	logger.Debug().Str("path", path).Msg("Ensuring directory exists")

	if _, err := os.Stat(path); os.IsNotExist(err) {
		logger.Debug().Str("path", path).Msg("Directory does not exist, creating it")
		if err := os.MkdirAll(path, 0755); err != nil {
			logger.Error().Err(err).Str("path", path).Msg("Error creating directory")
			return fmt.Errorf("error creating directory %s: %v", path, err)
		}
		logger.Debug().Str("path", path).Msg("Directory created successfully")
	} else {
		logger.Debug().Str("path", path).Msg("Directory already exists")
	}

	return nil
}

// CopyFile copies a file from source to destination
func CopyFile(src, dst string) error {
	logger := logging.GetLogger()
	logger.Debug().Str("src", src).Str("dst", dst).Msg("Copying file")

	// Ensure destination directory exists
	dstDir := filepath.Dir(dst)
	if err := EnsureDirExists(dstDir); err != nil {
		logger.Error().Err(err).Str("dir", dstDir).Msg("Error ensuring destination directory exists")
		return err
	}

	// Open source file
	srcFile, err := os.Open(src)
	if err != nil {
		logger.Error().Err(err).Str("src", src).Msg("Error opening source file")
		return fmt.Errorf("error opening source file %s: %v", src, err)
	}
	defer srcFile.Close()

	// Create destination file
	dstFile, err := os.Create(dst)
	if err != nil {
		logger.Error().Err(err).Str("dst", dst).Msg("Error creating destination file")
		return fmt.Errorf("error creating destination file %s: %v", dst, err)
	}
	defer dstFile.Close()

	// Copy contents
	bytesWritten, err := io.Copy(dstFile, srcFile)
	if err != nil {
		logger.Error().Err(err).Msg("Error copying file contents")
		return fmt.Errorf("error copying file contents: %v", err)
	}

	logger.Debug().Str("src", src).Str("dst", dst).Int64("bytes_written", bytesWritten).Msg("File copied successfully")
	return nil
}

// FileExists checks if a file exists
func FileExists(path string) bool {
	logger := logging.GetLogger()
	logger.Debug().Str("path", path).Msg("Checking if file exists")

	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		logger.Debug().Str("path", path).Msg("File does not exist")
		return false
	}

	isFile := !info.IsDir()
	logger.Debug().Str("path", path).Bool("exists", isFile).Msg("File existence check complete")
	return isFile
}

// DirExists checks if a directory exists
func DirExists(path string) bool {
	logger := logging.GetLogger()
	logger.Debug().Str("path", path).Msg("Checking if directory exists")

	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		logger.Debug().Str("path", path).Msg("Directory does not exist")
		return false
	}

	isDir := info.IsDir()
	logger.Debug().Str("path", path).Bool("exists", isDir).Msg("Directory existence check complete")
	return isDir
}
