package utf8s

import (
	"io"
)

// A utf8s.RuneReader implements the io.RuneReader interface by reading from an io.Reader.
type RuneReader struct {
	reader io.Reader
}

func NewRuneReader(reader io.Reader) *RuneReader {
	return &RuneReader{
		reader: reader,
	}
}

func (receiver *RuneReader) ReadRune() (rune, int, error) {
	reader := receiver.reader

	if nil == reader {
		return 0, 0, errNilReader
	}

	return ReadRune(reader)
}
