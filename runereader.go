package utf8

import (
	"io"
)

var _ io.RuneReader = WrapRuneReader(nil)
var _ io.RuneReader = NewRuneReader(nil)

// A utf8.RuneReader implements the io.RuneReader interface by reading from an io.Reader.
type RuneReader struct {
	reader io.Reader
}

func WrapRuneReader(reader io.Reader) RuneReader {
	return RuneReader{
		reader: reader,
	}
}

func NewRuneReader(reader io.Reader) *RuneReader {
	var runereader RuneReader = WrapRuneReader(reader)

	return &runereader
}

func (receiver RuneReader) ReadRune() (rune, int, error) {
	reader := receiver.reader
	if nil == reader {
		return 0, 0, errNilReader
	}

	return ReadRune(reader)
}
