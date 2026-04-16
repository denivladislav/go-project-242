package code

import (
	"fmt"

	"code/internal/calcsize"
	"code/internal/format"
)

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	var sizeOptions = calcsize.Options{
		Recursive: recursive,
		All:       all,
	}

	size, err := calcsize.CalcSize(path, sizeOptions)
	if err != nil {
		return "", fmt.Errorf("calc size failed: %w", err)
	}

	return format.FormatSize(size, human)
}
