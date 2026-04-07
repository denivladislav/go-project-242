package format

import (
	"fmt"
)

const denominator = 1024

var units = []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}

func convertSize(size int64) (float64, string) {
	curr := float64(size)

	for _, unit := range units[:len(units)-1] {
		new := curr / denominator

		if new < 1 {
			return curr, unit
		}

		curr = new
	}

	return curr, units[len(units)-1]
}

func formatByte(size int64) string {
	return fmt.Sprintf("%d%s", size, units[0])
}

func formatHuman(size float64, unit string) string {
	return fmt.Sprintf("%.1f%s", size, unit)
}

func FormatSize(size int64, human bool) (string, error) {
	if size < 0 {
		return "", fmt.Errorf("size cannot be < 0")
	}

	if !human {
		return formatByte(size), nil
	}

	convertedSize, unit := convertSize(size)

	if unit == units[0] {
		return formatByte(size), nil
	}

	return formatHuman(convertedSize, unit), nil
}
