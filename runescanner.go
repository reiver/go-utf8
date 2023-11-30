package utf8

import (
	"io"
)

var _ io.RuneScanner = NewRuneScanner(nil)

// A utf8.RuneScanner implements the io.RuneScanner interface by reading from an io.Reader.
type RuneScanner struct {
	reader io.Reader

	prevRune rune
	prevSize int
	prevErr  error

	peeked bool
}

func WrapRuneScanner(reader io.Reader) RuneScanner {
	return RuneScanner{
		reader: reader,
	}
}

func NewRuneScanner(reader io.Reader) *RuneScanner {
	var runescanner RuneScanner = WrapRuneScanner(reader)

	return &runescanner
}

// Buffered returns the number of bytes the UTF-8 encoding of the current buffered rune takes up, if there is a buffered rune.
//
// A buffered rune would come from someone calleding .UnreadRune().
//
// If there is not buffered rune then .Buffered() returns zero (0).
//
// So, for example, if .UnreadRune() was called for the rune 'A' (U+0041), then .Buffered() would return 1.
//
// Also, for example, if .UnreadRune() was called for the rune '۵' (U+06F5), then .Buffered() would return 2.
//
// And, for example, if .UnreadRune() was called for the rune '≡' (U+2261), then .Buffered() would return 3.
//
// And also, for example, if .UnreadRune() was called for the rune '🙂' (U+1F642), then .Buffered() would return 4.
//
// This method has been made to be semantically the same as bufio.Reader.Buffered()
func (receiver *RuneScanner) Buffered() int {
	if !receiver.peeked {
		return 0
	}

	return receiver.prevSize
}

func (receiver *RuneScanner) ReadRune() (rune, int, error) {
	if nil == receiver {
		return RuneError, 0, errNilReceiver
	}

	reader := receiver.reader
	if nil == reader {
		return RuneError, 0, errNilReader
	}

	if receiver.peeked {
		receiver.peeked = false
	} else {
		receiver.prevRune, receiver.prevSize, receiver.prevErr = ReadRune(reader)
	}


	return receiver.prevRune, receiver.prevSize, receiver.prevErr
}

func (receiver *RuneScanner) UnreadRune() error {
	if nil == receiver {
		return errNilReceiver
	}

	receiver.peeked = true

	return nil
}
