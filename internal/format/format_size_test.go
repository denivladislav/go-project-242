package format

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

type TestCase struct {
	desc     string
	size     int64
	human    bool
	expected string
}

func TestNegativeSize(t *testing.T) {
	testCase := TestCase{
		desc:  "negative",
		size:  -100,
		human: false,
	}

	result, err := FormatSize(testCase.size, testCase.human)
	require.Error(t, err)
	require.Equal(t, "", result)
}

func TestByteFormat(t *testing.T) {
	testCase := TestCase{
		desc:     "bytes",
		size:     100,
		human:    false,
		expected: "100B",
	}

	result, err := FormatSize(testCase.size, testCase.human)
	require.NoError(t, err)

	require.Equal(t, testCase.expected, result)
}

func TestHumanFormat(t *testing.T) {
	testCases := []TestCase{
		{
			desc:     "bytes",
			size:     100,
			human:    true,
			expected: "100B",
		},
		{
			desc:     "1 KB",
			size:     1024,
			human:    true,
			expected: "1.0KB",
		},
		{
			desc:     "1 KB with small overlap",
			size:     1024 + 1,
			human:    true,
			expected: "1.0KB",
		},
		{
			desc:     "1 KB with bigger overlap",
			size:     1024 + 100,
			human:    true,
			expected: "1.1KB",
		},
		{
			desc:     "MB",
			size:     int64(math.Pow(2, 20))*2 + 1024*400,
			human:    true,
			expected: "2.4MB",
		},
		{
			desc:     "EB",
			size:     int64(math.Pow(2, 60)),
			human:    true,
			expected: "1.0EB",
		},
		{
			desc:     "EB with overlap",
			size:     int64(math.Pow(2, 60) + math.Pow(2, 57)),
			human:    true,
			expected: "1.1EB",
		},
	}

	for _, testCase := range testCases {
		result, err := FormatSize(testCase.size, testCase.human)
		require.NoError(t, err)

		require.Equal(t, testCase.expected, result)
	}
}
