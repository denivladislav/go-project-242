package calcsize

import (
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

func calcDirSize(path string, options Options) (int64, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, fmt.Errorf("failed to read dir '%s': %w", path, err)
	}

	var size int64

	for _, entry := range entries {
		entryName := entry.Name()

		entryInfo, err := entry.Info()
		if err != nil {
			return 0, fmt.Errorf("failed to get entry info for '%s': %w", entryName, err)
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
			return 0, fmt.Errorf("calcDirSize failed: %w", err)
		}

		size += dirSize
	}

	return size, nil
}

func CalcSize(path string, options Options) (int64, error) {
	entry, err := os.Lstat(path)
	if err != nil {
		return 0, fmt.Errorf("failed to read '%s': %w", path, err)
	}

	if !options.All && isHidden(entry.Name()) {
		return 0, fmt.Errorf("no visible file or dir for '%s'", path)
	}

	if !entry.IsDir() {
		return entry.Size(), nil
	}

	dirSize, err := calcDirSize(path, options)
	if err != nil {
		return 0, fmt.Errorf("calcDirSize failed: %w", err)
	}

	return dirSize, nil
}
