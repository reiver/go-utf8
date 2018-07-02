package utf8s

import (
	"errors"
)

var (
	errInternalError = errors.New("Internal Error")
	errInvalidUTF8   = internalInvalidUTF8Complainer{}
	errNilReader     = internalNilReaderComplainer{}
	errNilWriter     = internalNilWriterComplainer{}
)
