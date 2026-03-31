package pathsize

import (
	"code/formatsize"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Human bool
	All   bool
}

func isHidden(name string) bool {
	return strings.HasPrefix(name, ".")
}

func getDirSize(path string, all bool) (int64, error) {
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

		if !all && isHidden(entry.Name()) {
			continue
		}

		// TODO: implement recursive check
		if !entryInfo.IsDir() {
			size += entryInfo.Size()
		}
	}

	return size, nil
}

func GetPathSize(path string, config Config) (string, error) {
	entry, err := os.Lstat(path)
	if err != nil {
		return "", err
	}

	if !config.All && isHidden(entry.Name()) {
		return "", fmt.Errorf("no visible file or dir with path %s", path)
	}

	var size int64

	if entry.IsDir() {
		dirSize, err := getDirSize(path, config.All)
		if err != nil {
			return "", err
		}

		size = dirSize
	} else {
		size = entry.Size()
	}

	return formatsize.FormatSize(size, config.Human)
}
