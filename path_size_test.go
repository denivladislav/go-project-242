package path_size

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnreachablePath(t *testing.T) {
	unreachablePath := "./unreachable"

	result, err := GetPathSize(unreachablePath)

	const msg = `GetPathSize should return "" and err for unreachable path`
	require.Error(t, err, msg)
	require.Equal(t, "", result, msg)
}

func TestEmptyFile(t *testing.T) {
	fixturePath := "./testdata/fileEmpty.txt"

	want := fmtPathSize(0)
	result, err := GetPathSize(fixturePath)

	require.NoError(t, err)
	require.Equal(t, want, result)
}

func TestFile(t *testing.T) {
	fixturePath := "./testdata/file.txt"

	entry, errLStat := os.Lstat(fixturePath)

	require.NoError(t, errLStat)

	want := fmtPathSize(entry.Size())
	result, errPathSize := GetPathSize(fixturePath)

	require.NoError(t, errPathSize)
	require.Equal(t, want, result)
}

func TestFolder(t *testing.T) {
	fixturePath := "./testData/testFolder"

	entry1, errLStat1 := os.Lstat(filepath.Join(fixturePath, "file1.txt"))
	entry2, errLStat2 := os.Lstat(filepath.Join(fixturePath, "file2.txt"))

	require.NoError(t, errLStat1)
	require.NoError(t, errLStat2)

	want := fmtPathSize(entry1.Size() + entry2.Size())
	result, errPathSize := GetPathSize(fixturePath)

	require.NoError(t, errPathSize)
	require.Equal(t, want, result)
}
