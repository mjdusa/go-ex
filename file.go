package ext

import (
	"os"
)

// FileSize returns the file size of the specified path file.
func FileSize(filePath string) (*int64, error) {
	file, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}

	size := file.Size()

	return &size, nil
}
