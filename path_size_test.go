package code

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

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

func validateError(t *testing.T, err error, checkErr func(err error) bool) {
	t.Helper()

	if !checkErr(err) {
		t.Errorf("error check failed for error: %v", err)
	}
}

func TestGetPathSizeUnreachable(t *testing.T) {
	path := filepath.Join(".", "unreachable")
	checkErr := func(err error) bool {
		return errors.Is(err, os.ErrNotExist)
	}

	_, err := GetPathSize(path, false, false, false)
	validateError(t, err, checkErr)
}

func TestGetPathSizeHappy(t *testing.T) {
	result, err := GetPathSize(testDataPath, true, true, true)
	assertError(t, err, false)

	assertEqual(t, result, "18B")
}
