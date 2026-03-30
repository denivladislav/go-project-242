package formatsize

import (
	"fmt"
)

const denominator = 1024

var units = []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}

func convertSize(size int64) (float64, string) {
	curr := float64(size)

	i := 0
	for i < len(units)-1 {
		new := curr / denominator
		if new < 1 {
			return curr, units[i]
		}

		curr = new
		i += 1
	}

	return curr, units[i]
}

func byteFormat(size int64) string {
	return fmt.Sprintf("%d%s", size, units[0])
}

func humanFormat(size float64, unit string) string {
	return fmt.Sprintf("%.1f%s", size, unit)
}

func FormatSize(size int64, human bool) (string, error) {
	if size < 0 {
		return "", fmt.Errorf("error: size cannot be < 0")
	}

	if human {
		convertedSize, unit := convertSize(size)

		if unit == units[0] {
			return byteFormat(size), nil
		}

		return humanFormat(convertedSize, unit), nil
	}

	return byteFormat(size), nil
}
