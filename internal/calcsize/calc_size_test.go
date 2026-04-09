package calcsize

import (
	"path/filepath"
	"testing"
)

type test struct {
	name     string
	path     string
	options  Options
	expected int64
	wantErr  bool
}

var actualSizes = map[string]int64{
	"emptyFile":                  0,
	"testDataFolder":             3,
	"testDataFolderRecursive":    11,
	"testDataFolderRecursiveAll": 18,
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

var testDataPath = filepath.Join("..", "..", "testdata")

func TestUnreachablePath(t *testing.T) {
	tt := test{
		name:    "returns error for unreachable path",
		path:    filepath.Join(".", "unreachable"),
		wantErr: true,
	}

	t.Run(tt.name, func(t *testing.T) {
		_, err := CalcSize(tt.path, tt.options)
		assertError(t, err, tt.wantErr)
	})
}

func TestSize(t *testing.T) {
	tests := []test{
		{
			name:     "returns 0 for empty file",
			path:     filepath.Join(testDataPath, "sizeEmpty.txt"),
			wantErr:  false,
			expected: actualSizes["emptyFile"],
		},
		{
			name:     "returns non-recursive folder size",
			path:     testDataPath,
			wantErr:  false,
			expected: actualSizes["testDataFolder"],
		},
		{
			name: "returns recursive folder size",
			path: testDataPath,
			options: Options{
				Recursive: true,
			},
			wantErr:  false,
			expected: actualSizes["testDataFolderRecursive"],
		},
		{
			name: "returns recursive + hidden folder size",
			path: testDataPath,
			options: Options{
				Recursive: true,
				All:       true,
			},
			wantErr:  false,
			expected: actualSizes["testDataFolderRecursiveAll"],
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := CalcSize(tt.path, tt.options)
			assertError(t, err, tt.wantErr)

			if tt.wantErr {
				return
			}

			assertEqual(t, result, tt.expected)
		})
	}
}
