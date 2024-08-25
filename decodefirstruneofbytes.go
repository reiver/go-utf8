package utf8

import (
	"io"
	"bytes"
)

// DecodeFirstRuneOfBytes decodes the first UTF-8 encoded rune at the beginning of the []byte.
func DecodeFirstRuneOfBytes(p []byte) (rune, int, error) {
	var reader io.Reader = bytes.NewReader(p)
	return ReadRune(reader)
}
