package utf8s

import (
	"io"
)

// WriteRune writes a single UTF-8 encoded Unicode character and returns the number of bytes written.
func WriteRune(writer io.Writer, r rune) (int, error) {
	if nil == writer {
		return 0, errNilWriter
	}

	switch {
	case 127 >= r:
		var buffer [1]byte

		buffer[0] = byte(r)

		p := buffer[:]

		return writer.Write(p)

	case 0x7FF >= r:
		var buffer [2]byte

		buffer[0] = 0xC0 | byte((0x000007C0 & r) >> 6)
		buffer[1] = 0x80 | byte( 0x0000003F & r)

		p := buffer[:]

		return writer.Write(p)

	case 0xFFFF >= r:
		var buffer [3]byte

		buffer[0] = 0xE0 | byte((0x0000F000 & r) >> 12)
		buffer[1] = 0x80 | byte((0x00000FC0 & r) >>  6)
		buffer[2] = 0x80 | byte( 0x0000003F & r)

		p := buffer[:]

		return writer.Write(p)

	case 0x10FFFF >= r:
		var buffer [4]byte

		buffer[0] = 0xF0 | byte((0x001C0000 & r) >> 18)
		buffer[1] = 0x80 | byte((0x0003F000 & r) >> 12)
		buffer[2] = 0x80 | byte((0x00000FC0 & r) >>  6)
		buffer[3] = 0x80 | byte( 0x0000003F & r)

		p := buffer[:]

		return writer.Write(p)

	default:
		return 0, errInternalError
	}

}
