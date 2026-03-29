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

func fmtPathSize(size int64) string {
	return fmt.Sprintf("%dB", size)
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

		return fmtPathSize(size), nil
	}

	return fmtPathSize(entry.Size()), nil
}
