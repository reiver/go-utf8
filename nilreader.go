package utf8

type NilReaderComplainer interface {
	error
	NilReaderComplainer()
}

type internalNilReaderComplainer struct{}

func (complainer internalNilReaderComplainer) Error() string {
	return "Nil Reader"
}

func (complainer internalNilReaderComplainer) NilReaderComplainer() {
	// Nothing here.
}
