package code

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

var byteConfig = Config{
	Human:     false,
	All:       false,
	Recursive: false,
}

func TestUnreachablePath(t *testing.T) {
	unreachablePath := "./unreachable"

	result, err := GetPathSize(
		unreachablePath,
		byteConfig.Recursive,
		byteConfig.Human,
		byteConfig.All,
	)

	const msg = `GetPathSize should return "" and err for unreachable path`
	require.Error(t, err, msg)
	require.Equal(t, "", result, msg)
}

func TestEmptyFile(t *testing.T) {
	fixturePath := "./testdata/fileEmpty.txt"

	want, err := FormatSize(0, false)
	require.NoError(t, err)

	result, err := GetPathSize(fixturePath, byteConfig.Recursive, byteConfig.Human, byteConfig.All)
	require.NoError(t, err)

	require.Equal(t, want, result)
}

func TestFile(t *testing.T) {
	fixturePath := "./testdata/file.txt"

	entry, err := os.Lstat(fixturePath)
	require.NoError(t, err)

	want, err := FormatSize(entry.Size(), false)
	require.NoError(t, err)

	result, err := GetPathSize(fixturePath, byteConfig.Recursive, byteConfig.Human, byteConfig.All)

	require.NoError(t, err)
	require.Equal(t, want, result)
}

func TestFolder(t *testing.T) {
	fixturePath := "./testdata/testFolder"

	entry1, err := os.Lstat(filepath.Join(fixturePath, "file1.txt"))
	require.NoError(t, err)

	entry2, err := os.Lstat(filepath.Join(fixturePath, "file2.txt"))
	require.NoError(t, err)

	want, err := FormatSize(entry1.Size()+entry2.Size(), false)
	require.NoError(t, err)

	result, err := GetPathSize(fixturePath, byteConfig.Recursive, byteConfig.Human, byteConfig.All)
	require.NoError(t, err)

	require.Equal(t, want, result)
}

func TestHidden(t *testing.T) {
	fixturePath := "./testdata/testFolder/.testFolderInnerHidden/.fileHidden.txt"

	entry, err := os.Lstat(fixturePath)
	require.NoError(t, err)

	want, err := FormatSize(entry.Size(), false)
	require.NoError(t, err)

	// Checking file is not tracked without --all flag
	_, err = GetPathSize(fixturePath, byteConfig.Recursive, byteConfig.Human, byteConfig.All)
	require.Error(t, err)

	// Checking file is tracked with --all flag
	result, err := GetPathSize(fixturePath, byteConfig.Recursive, byteConfig.Human, true)
	require.NoError(t, err)

	require.Equal(t, want, result)
}
