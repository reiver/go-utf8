package utf8

import (
	"bytes"
)

// FormatBinary returns a representation of a rune as a sequence of bytes, given in binary format.
//
// Example
//
//	utf8.FormatBinary('۵')
//	
//	// Outputs:
//	// <<0b11011011 ; 0b10110101>>
func FormatBinary(r rune) string {

	var buffer bytes.Buffer

	_, err := WriteRune(&buffer, r)
	if nil != err {
		return "<<ERROR>>"
	}


	var output bytes.Buffer

	length := buffer.Len()
	for i:=0; i<length; i++ {
		switch i {
		case 0:
			output.WriteString("<<0b")
		default:
			output.WriteString(" ; 0b")
		}


		b := buffer.Bytes()[i]

		switch {
		case 0 == (0x80 & b):
			output.WriteRune('0')
		default:
			output.WriteRune('1')
		}

		switch {
		case 0 == (0x40 & b):
			output.WriteRune('0')
		default:
			output.WriteRune('1')
		}

		switch {
		case 0 == (0x20 & b):
			output.WriteRune('0')
		default:
			output.WriteRune('1')
		}

		switch {
		case 0 == (0x10 & b):
			output.WriteRune('0')
		default:
			output.WriteRune('1')
		}

		switch {
		case 0 == (0x08 & b):
			output.WriteRune('0')
		default:
			output.WriteRune('1')
		}

		switch {
		case 0 == (0x04 & b):
			output.WriteRune('0')
		default:
			output.WriteRune('1')
		}

		switch {
		case 0 == (0x02 & b):
			output.WriteRune('0')
		default:
			output.WriteRune('1')
		}

		switch {
		case 0 == (0x01 & b):
			output.WriteRune('0')
		default:
			output.WriteRune('1')
		}
	}
	output.WriteString(">>")

	return output.String()
}
