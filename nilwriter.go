package utf8s

type NilWriterComplainer interface {
	error
	NilWriterComplainer()
}

type internalNilWriterComplainer struct{}

func (complainer internalNilWriterComplainer) Error() string {
	return "Nil Writer"
}

func (complainer internalNilWriterComplainer) NilWriterComplainer() {
	// Nothing here.
}
