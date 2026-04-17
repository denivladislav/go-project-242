package format

import (
	"errors"
	"testing"
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
)

func assertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()

	if got != want {
		t.Errorf("got = %v, want = %v", got, want)
	}
}

func assertError(t *testing.T, err error, wantErr bool) {
	t.Helper()

	if (err != nil) != wantErr {
		t.Errorf("error = %v, wantErr = %v", err, wantErr)
	}
}

func validateError(t *testing.T, err error, checkErr func(err error) bool) {
	t.Helper()

	if !checkErr(err) {
		t.Errorf("error check failed for error: %v", err)
	}
}

func TestNegativeSize(t *testing.T) {
	checkErr := func(err error) bool {
		return errors.Is(err, ErrNegativeSize)
	}

	_, err := FormatSize(-MB, false)
	validateError(t, err, checkErr)
}

func TestByteFormat(t *testing.T) {
	result, err := FormatSize(10*KB, false)
	assertError(t, err, false)
	assertEqual(t, result, "10240B")
}

func TestHumanFormat(t *testing.T) {
	type test struct {
		name     string
		size     int64
		expected string
	}

	tests := []test{
		{
			name:     "keeps bytes for small size",
			size:     100,
			expected: "100B",
		},
		{
			name:     "rounds 1024B to 1.0KB",
			size:     1024,
			expected: "1.0KB",
		},
		{
			name:     "rounds 1KB + 1B to 1.0KB",
			size:     KB + 1,
			expected: "1.0KB",
		},
		{
			name:     "normalizes 1KB + 100B to 1.1KB",
			size:     KB + 100,
			expected: "1.1KB",
		},
		{
			name:     "normalizes 2MB + 400KB to 2.4MB",
			size:     2*MB + 400*KB,
			expected: "2.4MB",
		},
		{
			name:     "rounds 1024PB to 1.0EB",
			size:     1024 * PB,
			expected: "1.0EB",
		},
		{
			name:     "normalizes 1EB + 100PB to 1.1EB",
			size:     EB + 100*PB,
			expected: "1.1EB",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := FormatSize(tt.size, true)
			assertError(t, err, false)

			assertEqual(t, result, tt.expected)
		})
	}
}
