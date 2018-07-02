package utf8s

// Len returns the number of bytes in a UTF-8 encoding of this Unicode code point.
//
// Example
//
//	length := utf8s.Len('A')
//
//	// length == 1
//
// Example
//
//	length := utf8s.Len('r')
//
//	// length == 1
//
// Example
//
//	length := utf8s.Len('¡')
//
//	// length == 2
//
// Example
//
//	length := utf8s.Len('۵')
//
//	// length == 2
func Len(r rune) int {

	switch {
	case 127 >= r:
		return 1

	case 0x7FF >= r:
		return 2

	case 0xFFFF >= r:
		return 3

	case 0x10FFFF >= r:
		return 4

	default:
		return 0
	}
}
