package utf8

// InvalidUTF8Complainer is a type of error that could be returned
// by the utf8.ReadRune() function,
// by the utf8.RuneReader.ReadRune() method, and
// by the utf8.RuneScanner.ReadRune() method.
//
// Here is how one might use this type:
//
//	r, n, err := utf8.ReadRune(reader)
//	if nil != err {
//		switch {
//		case utf8.InvalidUTF8Complainer:
//			//@TODO
//		default:
//			//@TODO
//		}
//	}
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
