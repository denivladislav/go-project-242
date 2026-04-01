package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	Human     bool
	All       bool
	Recursive bool
}

func isHidden(name string) bool {
	return strings.HasPrefix(name, ".")
}

func getDirSize(path string, all, recursive bool) (int64, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}

	var size int64
	for _, entry := range entries {
		entryInfo, err := entry.Info()
		if err != nil {
			return 0, err
		}

		entryName := entryInfo.Name()

		if !all && isHidden(entryName) {
			continue
		}

		if entryInfo.IsDir() {
			if !recursive {
				continue
			}

			dirSize, err := getDirSize(filepath.Join(path, entryName), all, recursive)
			if err != nil {
				return 0, err
			}

			size += dirSize
		} else {
			size += entryInfo.Size()
		}
	}

	return size, nil
}

func GetPathSize(path string, human, all, recursive bool) (string, error) {
	entry, err := os.Lstat(path)
	if err != nil {
		return "", err
	}

	if !all && isHidden(entry.Name()) {
		return "", fmt.Errorf("no visible file or dir with path %s", path)
	}

	var size int64

	if entry.IsDir() {
		dirSize, err := getDirSize(path, all, recursive)
		if err != nil {
			return "", err
		}

		size = dirSize
	} else {
		size = entry.Size()
	}

	return FormatSize(size, human)
}
