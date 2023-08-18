package utf8

import (
	"errors"
)

var (
	errInternalError = errors.New("Internal Error")
	errInvalidUTF8   = internalInvalidUTF8Error{}
	errNilReader     = internalNilReaderError{}
	errNilReceiver   = errors.New("Nil Receiver")
	errNilWriter     = internalNilWriterError{}
)
