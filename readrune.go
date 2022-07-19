package utf8

import (
	"io"
)

// ReadRune reads a single UTF-8 encoded Unicode character from an io.Reader,
// and returns the Unicode character (as a Go rune) and the number of bytes read.
//
// If ‘reader’ is nil then ReaderRune will return an error that matches utf8.NilReaderComplainer.
//
// Example
//
// Here is an example usage of ReadRune:
//
//      r, n, err := utf8.ReadRune(reader)
//      if nil != err {
//
//              switch err.(type) {
//              case utf8.NilReaderComplainer:
//                      //@TODO
//              case utf8.InvalidUTF8Complainer:
//                      //@TODO
//              default:
//                      //TODO
//              }
//
//      }
//
// Number Of Bytes
//
// Note that a single UTF-8 encoded Unicode character could be more than one byte.
//
// For example, the Unicode "≡" (IDENTICAL TO) character gets encoded using 3 bytes under UTF-8.
func ReadRune(reader io.Reader) (rune, int, error) {
	if nil == reader {
		return 0, 0, errNilReader
	}

	var count int

	var b0 byte
	{
		var buffer [1]byte
		var p []byte = buffer[:]

		n, err := reader.Read(p)
		count += n
		if nil != err {
			return 0, count, err
		}
		if 1 != n {
			return 0, count, errInternalError
		}

		b0 = buffer[0]
	}

	if 127 >= b0 {
		return rune(b0), count, nil
	}

	var more int
	{
		switch {

		//   110x,xxxx       110x,xxxx
		// 0b1100,0000 == (0b1110,0000 & b0)
		case 0xC0 == (0xE0 & b0):
			more = 2-1

		//   1110,xxxx       1110,xxxx
		// 0b1110,0000 == (0b1111,0000 & b0)
		case 0xE0 == (0xF0 & b0):
			more = 3-1

		//   1111,0xxx       1111,0xxx
		// 0b1111,0000 == (0b1111,1000 & b0)
		case 0xF0 == (0xF8 & b0):
			more = 4-1

		//   1111,10xx       1111,10xx
		// 0b1111,1000 == (0b1111,1100 & b0)
		case 0xF8 == (0xFC & b0):
			more = 5-1

		//   1111,110x       1111,110x
		// 0b1111,1100 == (0b1111,1110 & b0)
		case 0xFC == (0xFE & b0):
			more = 6-1

		//   1111,1111       1111,1111
		// 0b1111,1110 == (0b1111,1111 & b0)
		case 0xFE == (0xFF & b0):
			more = 7-1

		default:
			return 0, count, errInternalError
		}
	}


	var bs [6]byte
	{
		p := bs[:more]

		n, err := reader.Read(p)
		count += n
		if nil != err {
			return 0, count, err
		}
		if more != n {
			return 0, count, errInternalError
		}
	}

	var r rune
	{

		var b byte

		switch {

		//   110x,xxxx       110x,xxxx
		// 0b1100,0000 == (0b1110,0000 & b0)
		case 0xC0 == (0xE0 & b0):
			b = (0xE0^0xFF) & b0

		//   1110,xxxx       1110,xxxx
		// 0b1110,0000 == (0b1111,0000 & b0)
		case 0xE0 == (0xF0 & b0):
			b = (0xF0^0xFF) & b0

		//   1111,0xxx       1111,0xxx
		// 0b1111,0000 == (0b1111,1000 & b0)
		case 0xF0 == (0xF8 & b0):
			b = (0xF8^0xFF) & b0

		//   1111,10xx       1111,10xx
		// 0b1111,1000 == (0b1111,1100 & b0)
		case 0xF8 == (0xFC & b0):
			b = (0xFC^0xFF) & b0

		//   1111,110x       1111,110x
		// 0b1111,1100 == (0b1111,1110 & b0)
		case 0xFC == (0xFE & b0):
			b = (0xFE^0xFF) & b0

		//   1111,1111       1111,1111
		// 0b1111,1110 == (0b1111,1111 & b0)
		case 0xFE == (0xFF & b0):
			//b := (0xFF^0xFF) & b0

		default:
			return 0, count, errInternalError
		}

		r = rune(b)
		r <<= 6

		for i:=0; i<more; i++ {

			bsi := bs[i]

			// if 0b1000,0000 != (0b0b1100,0000 & bsi) {
			if 0x80 != (0xC0 & bsi) {
				return 0, count, errInvalidUTF8
			}


			//  b := 0b0011,1111 & bsi
			b := 0x3F & bsi

			r2 := rune(b)

			r |= r2
			if i < (more-1) {
				r <<= 6
			}
		}
	}

	return r, count, nil
}
