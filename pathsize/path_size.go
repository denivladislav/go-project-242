package pathsize

import (
	"code/formatsize"
	"os"
)

func getDirSize(path string) (int64, error) {
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

		// TODO: implement recursive check
		if !entryInfo.IsDir() {
			size += entryInfo.Size()
		}
	}

	return size, nil
}

func GetPathSize(path string, human bool) (string, error) {
	entry, err := os.Lstat(path)
	if err != nil {
		return "", err
	}

	if entry.IsDir() {
		size, err := getDirSize(path)
		if err != nil {
			return "", err
		}

		return formatsize.FormatSize(size, human)
	}

	return formatsize.FormatSize(entry.Size(), human)
}
