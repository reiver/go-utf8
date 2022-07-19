package utf8

import (
	"io"
)

// A utf8.RuneReader implements the io.RuneReader interface by reading from an io.Reader.
type RuneReader struct {
	reader io.Reader
}

func RuneReaderWrap(reader io.Reader) RuneReader {
	return RuneReader{
		reader: reader,
	}
}

func (receiver *RuneReader) ReadRune() (rune, int, error) {
	if nil == receiver {
		return 0, 0, errNilReceiver
	}

	reader := receiver.reader
	if nil == reader {
		return 0, 0, errNilReader
	}

	return ReadRune(reader)
}
