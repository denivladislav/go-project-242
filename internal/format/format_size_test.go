package format

import (
	"math"
	"testing"
)

type test struct {
	name     string
	size     int64
	human    bool
	expected string
	wantErr  bool
}

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

func TestNegativeSize(t *testing.T) {
	tt := test{
		name:     "returns error for negative size",
		size:     -100,
		human:    false,
		expected: "",
		wantErr:  true,
	}

	t.Run(tt.name, func(t *testing.T) {
		_, err := FormatSize(tt.size, tt.human)
		assertError(t, err, tt.wantErr)
	})
}

func TestByteFormat(t *testing.T) {
	tt := test{
		name:     "returns bytes for size",
		size:     10000,
		human:    false,
		expected: "10000B",
		wantErr:  false,
	}

	t.Run(tt.name, func(t *testing.T) {
		result, err := FormatSize(tt.size, tt.human)
		assertError(t, err, tt.wantErr)
		assertEqual(t, result, tt.expected)
	})
}

func TestHumanFormat(t *testing.T) {
	tests := []test{
		{
			name:     "returns bytes for size < 1024B",
			size:     100,
			human:    true,
			expected: "100B",
		},
		{
			name:     "returns 1.0KB for 1024B",
			size:     1024,
			human:    true,
			expected: "1.0KB",
		},
		{
			name:     "returns 1.0KB for 1KB + 1B",
			size:     1024 + 1,
			human:    true,
			expected: "1.0KB",
		},
		{
			name:     "returns 1.1KB for 1KB + 100B",
			size:     1024 + 100,
			human:    true,
			expected: "1.1KB",
		},
		{
			name:     "returns 2.4MB for 2MB + 400KB",
			size:     int64(math.Pow(2, 20))*2 + 1024*400,
			human:    true,
			expected: "2.4MB",
		},
		{
			name:     "returns 1.0EB for 2^60B",
			size:     int64(math.Pow(2, 60)),
			human:    true,
			expected: "1.0EB",
		},
		{
			name:     "returns 1.1EB for 1EB + 100PB",
			size:     int64(math.Pow(2, 60) + math.Pow(2, 50)*100),
			human:    true,
			expected: "1.1EB",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := FormatSize(tt.size, tt.human)
			assertError(t, err, tt.wantErr)

			if tt.wantErr {
				return
			}

			assertEqual(t, result, tt.expected)
		})
	}
}
