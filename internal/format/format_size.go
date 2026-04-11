package format

import (
	"fmt"
)

const denominator = 1024

var units = []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}

func normalizeSize(size int64) (float64, string) {
	value := float64(size)
	index := 0

	for value >= denominator && index < len(units)-1 {
		value = value / denominator
		index += 1
	}

	return value, units[index]
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

	normalizedSize, unit := normalizeSize(size)

	if unit == units[0] {
		return formatByte(size), nil
	}

	return formatHuman(normalizedSize, unit), nil
}
