package fs

import (
	"crypto/sha1"
	"io"
	"os"
)

// GetFileSHA1 get file sha1
func GetFileSHA1(path string) ([]byte, error) {
	f, fErr := os.Open(path)
	if fErr != nil {
		return nil, fErr
	}
	defer f.Close()

	h := sha1.New()
	if _, err := io.Copy(h, f); err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}
