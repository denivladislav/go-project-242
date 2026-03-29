package path_size

import (
	"fmt"
	"os"
)

func getDirSize(path string) (int64, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}

	var size int64
	for _, file := range files {
		fileInfo, err := file.Info()
		if err != nil {
			return 0, err
		}

		size += fileInfo.Size()
	}

	return size, nil
}

func GetPathSize(path string) (string, error) {
	entry, err := os.Lstat(path)
	if err != nil {
		return "", err
	}

	if entry.IsDir() {
		size, err := getDirSize(path)
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("%dB", size), nil
	}

	return fmt.Sprintf("%dB", entry.Size()), nil
}
