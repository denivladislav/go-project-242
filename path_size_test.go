package code

import (
	"path/filepath"
	"testing"
)

type test struct {
	name     string
	path     string
	expected string
	wantErr  bool
}

var testDataPath = filepath.Join(".", "testdata")

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

func TestGetPathSizeFail(t *testing.T) {
	tt := test{
		name:    "returns error for unreachable path",
		path:    filepath.Join(".", "unreachable"),
		wantErr: true,
	}

	t.Run(tt.name, func(t *testing.T) {
		_, err := GetPathSize(tt.path, false, false, false)
		assertError(t, err, tt.wantErr)
	})
}

func TestGetPathSizeHappy(t *testing.T) {
	tt := test{
		name:     "returns all recursive folder size",
		path:     testDataPath,
		wantErr:  false,
		expected: "18B",
	}

	t.Run(tt.name, func(t *testing.T) {
		result, err := GetPathSize(tt.path, true, true, true)
		assertError(t, err, tt.wantErr)

		assertEqual(t, result, tt.expected)
	})
}
