package utf8

import (
	"io"
)

// A utf8.RuneScanner implements the io.RuneScanner interface by reading from an io.Reader.
type RuneScanner struct {
	reader io.Reader

	prevRune rune
	prevSize int
	prevErr  error

	peeked bool
}

func RuneScannerWrap(reader io.Reader) RuneScanner {
	return RuneScanner{
		reader: reader,
	}
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
