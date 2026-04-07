package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"code/internal/format"
)

type dirSizeOptions struct {
	all       bool
	recursive bool
}
type PathSizeOptions struct {
	All       bool
	Recursive bool
	Human     bool
}

func isHidden(name string) bool {
	return strings.HasPrefix(name, ".")
}

func getDirSize(path string, options dirSizeOptions) (int64, error) {
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

		if !options.all && isHidden(entryName) {
			continue
		}

		// isFile
		if !entryInfo.IsDir() {
			size += entryInfo.Size()
			continue
		}

		if !options.recursive {
			continue
		}

		newFilepath := filepath.Join(path, entryName)
		dirSize, err := getDirSize(newFilepath, options)
		if err != nil {
			return 0, err
		}

		size += dirSize
	}

	return size, nil
}

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	entry, err := os.Lstat(path)
	if err != nil {
		return "", err
	}

	if !all && isHidden(entry.Name()) {
		return "", fmt.Errorf("no visible file or dir with path %s", path)
	}

	var size int64

	// isFile
	if !entry.IsDir() {
		size = entry.Size()
		return format.FormatSize(size, human)
	}

	dirSizeOptions := dirSizeOptions{
		all:       all,
		recursive: recursive,
	}
	dirSize, err := getDirSize(path, dirSizeOptions)
	if err != nil {
		return "", err
	}

	size = dirSize

	return format.FormatSize(size, human)
}
