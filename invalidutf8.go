package utf8

// InvalidUTF8Error is a type of error that could be returned
// by the utf8.ReadRune() function,
// by the utf8.RuneReader.ReadRune() method, and
// by the utf8.RuneScanner.ReadRune() method.
//
// Here is how one might use this type:
//
//	r, n, err := utf8.ReadRune(reader)
//	if nil != err {
//		switch {
//		case utf8.InvalidUTF8Error:
//			//@TODO
//		default:
//			//@TODO
//		}
//	}
type InvalidUTF8Error interface {
	error
	InvalidUTF8Error()
}

type internalInvalidUTF8Error struct{}

func (complainer internalInvalidUTF8Error) Error() string {
	return "Invalid UTF-8"
}

func (complainer internalInvalidUTF8Error) InvalidUTF8Error() {
	// Nothing here.
}
