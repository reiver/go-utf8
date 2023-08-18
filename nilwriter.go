package utf8

type NilWriterError interface {
	error
	NilWriterError()
}

type internalNilWriterError struct{}

func (complainer internalNilWriterError) Error() string {
	return "Nil Writer"
}

func (complainer internalNilWriterError) NilWriterError() {
	// Nothing here.
}
