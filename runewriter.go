package utf8

import (
	"io"
)

// RuneWriter writes a single UTF-8 encoded Unicode characters.
type RuneWriter struct {
	writer io.Writer
}

// RuneWriterWrap wraps an io.Writer and returns a RuneWriter.
func RuneWriterWrap(writer io.Writer) RuneWriter {
	return RuneWriter{
		writer: writer,
	}
}

// WriteRune writes a single UTF-8 encoded Unicode character and returns the number of bytes written.
func (receiver *RuneWriter) WriteRune(r rune) (int, error) {
	if nil == receiver {
		return 0, errNilReceiver
	}

	return WriteRune(receiver.writer, r)
}
