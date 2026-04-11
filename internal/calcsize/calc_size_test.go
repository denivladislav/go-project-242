package calcsize

import (
	"path/filepath"
	"testing"
)

type test struct {
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
		path:    filepath.Join(".", "unreachable"),
		wantErr: true,
	}

	t.Run("unreachable path", func(t *testing.T) {
		_, err := CalcSize(tt.path, tt.options)
		assertError(t, err, tt.wantErr)
	})
}

func TestSize(t *testing.T) {
	tests := map[string]test{
		"empty file": {
			path:     filepath.Join(testDataPath, "sizeEmpty.txt"),
			wantErr:  false,
			expected: actualSizes["emptyFile"],
		},
		"non-recursive folder": {
			path:     testDataPath,
			wantErr:  false,
			expected: actualSizes["testDataFolder"],
		},
		"recursive folder": {
			path: testDataPath,
			options: Options{
				Recursive: true,
			},
			wantErr:  false,
			expected: actualSizes["testDataFolderRecursive"],
		},
		"recursive + hidden folder": {
			path: testDataPath,
			options: Options{
				Recursive: true,
				All:       true,
			},
			wantErr:  false,
			expected: actualSizes["testDataFolderRecursiveAll"],
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			result, err := CalcSize(tt.path, tt.options)
			assertError(t, err, tt.wantErr)

			if tt.wantErr {
				return
			}

			assertEqual(t, result, tt.expected)
		})
	}
}
