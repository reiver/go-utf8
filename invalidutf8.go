package utf8s

type InvalidUTF8Complainer interface {
	error
	InvalidUTF8Complainer()
}

type internalInvalidUTF8Complainer struct{}

func (complainer internalInvalidUTF8Complainer) Error() string {
	return "Invalid UTF-8"
}

func (complainer internalInvalidUTF8Complainer) InvalidUTF8Complainer() {
	// Nothing here.
}
