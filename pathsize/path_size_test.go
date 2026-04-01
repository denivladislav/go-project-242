package pathsize

import (
	"code/formatsize"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

var byteConfig = Config{
	Human: false,
	All:   false,
}

func TestUnreachablePath(t *testing.T) {
	unreachablePath := "./unreachable"

	result, err := GetPathSize(unreachablePath, byteConfig)

	const msg = `GetPathSize should return "" and err for unreachable path`
	require.Error(t, err, msg)
	require.Equal(t, "", result, msg)
}

func TestEmptyFile(t *testing.T) {
	fixturePath := "../testdata/fileEmpty.txt"

	want, err := formatsize.FormatSize(0, false)
	require.NoError(t, err)

	result, err := GetPathSize(fixturePath, byteConfig)
	require.NoError(t, err)

	require.Equal(t, want, result)
}

func TestFile(t *testing.T) {
	fixturePath := "../testdata/file.txt"

	entry, err := os.Lstat(fixturePath)
	require.NoError(t, err)

	want, err := formatsize.FormatSize(entry.Size(), false)
	require.NoError(t, err)

	result, err := GetPathSize(fixturePath, byteConfig)

	require.NoError(t, err)
	require.Equal(t, want, result)
}

func TestFolder(t *testing.T) {
	fixturePath := "../testData/testFolder"

	entry1, err := os.Lstat(filepath.Join(fixturePath, "file1.txt"))
	require.NoError(t, err)

	entry2, err := os.Lstat(filepath.Join(fixturePath, "file2.txt"))
	require.NoError(t, err)

	want, err := formatsize.FormatSize(entry1.Size()+entry2.Size(), false)
	require.NoError(t, err)

	result, err := GetPathSize(fixturePath, byteConfig)
	require.NoError(t, err)

	require.Equal(t, want, result)
}

func TestHidden(t *testing.T) {
	fixturePath := "../testData/testFolder/.testFolderInnerHidden/.fileHidden.txt"

	entry, err := os.Lstat(fixturePath)
	require.NoError(t, err)

	want, err := formatsize.FormatSize(entry.Size(), false)
	require.NoError(t, err)

	// Checking file is not tracked without --all flag
	_, err = GetPathSize(fixturePath, byteConfig)
	require.Error(t, err)

	// Checking file is tracked with --all flag
	result, err := GetPathSize(fixturePath, Config{
		Human: true,
		All:   true,
	})
	require.NoError(t, err)

	require.Equal(t, want, result)
}
