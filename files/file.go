package files

import (
	"os"
	"testing"
)

func MustOpen(t *testing.T, path string) *os.File {
	t.Helper()
	f, err := os.Open(path)
	if err != nil {
		t.Fatal(err)
	}

	return f
}
