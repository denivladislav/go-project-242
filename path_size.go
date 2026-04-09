package code

import (
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
		return "", err
	}

	return format.FormatSize(size, human)
}
