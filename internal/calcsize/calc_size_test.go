package calcsize

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
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

var testDataPath = filepath.Join("..", "..", "testdata")

func TestCalcSizeErrors(t *testing.T) {
	type test struct {
		path     string
		checkErr func(err error) bool
	}

	tests := map[string]test{
		"unreachable path causes error": {
			path: filepath.Join(".", "unreachable"),
			checkErr: func(err error) bool {
				return errors.Is(err, os.ErrNotExist)
			},
		},
		"no visible entry causes error": {
			path: filepath.Join(testDataPath, "nested", ".innerHidden"),
			checkErr: func(err error) bool {
				return errors.Is(err, ErrNoVisibleEntry)
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := CalcSize(tt.path, Options{})
			validateError(t, err, tt.checkErr)
		})
	}
}

func TestCalcSize(t *testing.T) {
	type test struct {
		path     string
		options  Options
		expected int64
	}

	tests := map[string]test{
		"calcs empty file size as 0B": {
			path:     filepath.Join(testDataPath, "sizeEmpty.txt"),
			expected: 0,
		},
		"calcs non-recursive folder size": {
			path:     testDataPath,
			expected: 3,
		},
		"calcs recursive folder size": {
			path: testDataPath,
			options: Options{
				Recursive: true,
			},
			expected: 11,
		},
		"calcs recursive + hidden folder size": {
			path: testDataPath,
			options: Options{
				Recursive: true,
				All:       true,
			},
			expected: 18,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			result, err := CalcSize(tt.path, tt.options)
			assertError(t, err, false)

			assertEqual(t, result, tt.expected)
		})
	}
}
