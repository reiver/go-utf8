package utf8

type NilReaderError interface {
	error
	NilReaderError()
}

type internalNilReaderError struct{}

func (complainer internalNilReaderError) Error() string {
	return "Nil Reader"
}

func (complainer internalNilReaderError) NilReaderError() {
	// Nothing here.
}
