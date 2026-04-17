package calcsize

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Options struct {
	Recursive bool
	All       bool
}

func isHidden(name string) bool {
	return strings.HasPrefix(name, ".")
}

var ErrNoVisibleEntry = errors.New("no visible file or dir")

func calcDirSize(path string, options Options) (int64, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, fmt.Errorf("os read dir failed: %w", err)
	}

	var size int64

	for _, entry := range entries {
		entryName := entry.Name()

		entryInfo, err := entry.Info()
		if err != nil {
			return 0, fmt.Errorf("entry info failed: %w", err)
		}

		if !options.All && isHidden(entryName) {
			continue
		}

		if !entryInfo.IsDir() {
			size += entryInfo.Size()
			continue
		}

		if !options.Recursive {
			continue
		}

		newFilepath := filepath.Join(path, entryName)

		dirSize, err := calcDirSize(newFilepath, options)
		if err != nil {
			return 0, fmt.Errorf("calc dir size failed: %w", err)
		}

		size += dirSize
	}

	return size, nil
}

func CalcSize(path string, options Options) (int64, error) {
	entry, err := os.Lstat(path)
	if err != nil {
		return 0, fmt.Errorf("os lstat failed: %w", err)
	}

	if !options.All && isHidden(entry.Name()) {
		return 0, fmt.Errorf("calc size failed for path '%s': %w", path, ErrNoVisibleEntry)
	}

	if !entry.IsDir() {
		return entry.Size(), nil
	}

	dirSize, err := calcDirSize(path, options)
	if err != nil {
		return 0, fmt.Errorf("calc dir size failed: %w", err)
	}

	return dirSize, nil
}
