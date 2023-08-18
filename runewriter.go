package utf8

import (
	"io"
)

type runeWriter interface {
	WriteRune(rune) (int, error)
}

var _ runeWriter = WrapRuneWriter(nil)

// RuneWriter writes a single UTF-8 encoded Unicode characters.
type RuneWriter struct {
	writer io.Writer
}

// WrapRuneWriter wraps an io.Writer and returns a RuneWriter.
func WrapRuneWriter(writer io.Writer) RuneWriter {
	return RuneWriter{
		writer: writer,
	}
}

// WriteRune writes a single UTF-8 encoded Unicode character and returns the number of bytes written.
func (receiver RuneWriter) WriteRune(r rune) (int, error) {
	return WriteRune(receiver.writer, r)
}
