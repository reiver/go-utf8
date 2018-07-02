package utf8s

import (
	"bytes"

	"testing"
)

func TestWriteRune(t *testing.T) {

	tests := []struct{
		Rune            rune
		ExpectedInt     int
		ExpectedBytes []byte
	}{
		{
			Rune:         'a',
			ExpectedInt:   1,
			ExpectedBytes: []byte{'a'},
		},



		{
			Rune:         'b',
			ExpectedInt:   1,
			ExpectedBytes: []byte{'b'},
		},



		{
			Rune:         'c',
			ExpectedInt:   1,
			ExpectedBytes: []byte{'c'},
		},



		{
			Rune:         'A',
			ExpectedInt:   1,
			ExpectedBytes: []byte{'A'},
		},



		{
			Rune:         'r',
			ExpectedInt:   1,
			ExpectedBytes: []byte{'r'},
		},



		{
			Rune:         '¡',
			ExpectedInt:   2,
			ExpectedBytes: []byte{0xC2, 0xA1},
		},



		{
			Rune:         '۵',
			ExpectedInt:   2,
			ExpectedBytes: []byte{0xDB, 0xB5},
		},



		{
			Rune:         '‱',
			ExpectedInt:   3,
			ExpectedBytes: []byte{0xe2, 0x80, 0xb1},
		},



		{
			Rune:         '≡',
			ExpectedInt:   3,
			ExpectedBytes: []byte{0xE2, 0x89, 0xA1},
		},



		{
			Rune:         '𐏕',
			ExpectedInt:   4,
			ExpectedBytes: []byte{0xf0, 0x90, 0x8f, 0x95},
		},



		{
			Rune:         '🙂',
			ExpectedInt:   4,
			ExpectedBytes: []byte{0xf0, 0x9f, 0x99, 0x82},
		},



		{
			Rune        :  0x0000,
			ExpectedInt:   1,
			ExpectedBytes: []byte{0x00},
		},
		{
			Rune:          0x0001,
			ExpectedInt:   1,
			ExpectedBytes: []byte{0x01},
		},
		{
			Rune:          0x007e,
			ExpectedInt:   1,
			ExpectedBytes: []byte{0x7e},
		},
		{
			Rune:          0x007f,
			ExpectedInt:   1,
			ExpectedBytes: []byte{0x7f},
		},



		{
			Rune:          0x0080,  // 0b0000,1000,0000
			ExpectedInt:   2,
			ExpectedBytes: []byte{0xC2, 0x80}, // <<0b11000010 ; 0b1000,0000>>
		},
		{
			Rune:          0x0081,  // 0b0000,1000,0000
			ExpectedInt:   2,
			ExpectedBytes: []byte{0xC2, 0x81}, // <<0b11000010 ; 0b1000,0001>>
		},
		{
			Rune:          0x07fe, // 0b0111,1111,1110
			ExpectedInt:   2,
			ExpectedBytes: []byte{0xDF, 0xBE}, // <<0b1101,1111 ; 0b1011,1110>>
		},
		{
			Rune:          0x07ff, // 0b0111,1111,1111
			ExpectedInt:   2,
			ExpectedBytes: []byte{0xDF, 0xBF}, // <<0b1101,1111 ; 0b1011,1111>>
		},



		{
			Rune:          0x0800, // 0b1000,0000,0000
			ExpectedInt:   3,
			ExpectedBytes: []byte{0xe0, 0xa0, 0x80}, // <<0b111,00000 ; 0b1010,0000 ; 0b1000,0000>>
		},
		{
			Rune:          0x0801, // 0b1000,0000,0001
			ExpectedInt:   3,
			ExpectedBytes: []byte{0xe0, 0xa0, 0x81}, // <<0b111,00000 ; 0b1010,0000 ; 0b1000,0001>>
		},
		{
			Rune:          0xfffe, // 0b1111,1111,1111,1110
			ExpectedInt:   3,
			ExpectedBytes: []byte{0xEF, 0xBF, 0xBE}, // <<0b11101111 ; 0b1011,1111 ; 0b1011,1110>>
		},
		{
			Rune:          0xffff, // 0b1111,1111,1111,1111
			ExpectedInt:   3,
			ExpectedBytes: []byte{0xEF, 0xBF, 0xBF}, // <<0b11101111 ; 0b1011,1111 ; 0b1011,1111>>
		},



		{
			Rune:          0x010000, // 0b0001,0000,0000,0000,0000
			ExpectedInt:   4,
			ExpectedBytes: []byte{0xF0, 0x90, 0x80, 0x80}, // <<0b1111,0000 ; 0b10010000 ; 0b1000,0000 ; 0b1000,0000>>
		},
		{
			Rune:          0x010001, // 0b0001,0000,0000,0000,0001
			ExpectedInt:   4,
			ExpectedBytes: []byte{0xF0, 0x90, 0x80, 0x81}, // <<0b1111,0000 ; 0b10010000 ; 0b1000,0000 ; 0b1000,0001>>
		},
		{
			Rune:          0x10fffe, // 0b0001,0000,1111,1111,1111,1110
			ExpectedInt:   4,
			ExpectedBytes: []byte{0xF4, 0x8F, 0xBF, 0xBE}, // <<0b1111,0100 ; 0b10001111 ; 0b1011,1111 ; 0b1011,1110>>
		},
		{
			Rune:          0x10ffff, // 0b0001,0000,1111,1111,1111,1111
			ExpectedInt:   4,
			ExpectedBytes: []byte{0xF4, 0x8F, 0xBF, 0xBF}, // <<0b1111,0100 ; 0b10001111 ; 0b1011,1111 ; 0b1011,1111>>
		},
	}


	for testNumber, test := range tests {

		var buffer bytes.Buffer

		actualInt, err := WriteRune(&buffer, test.Rune)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q", testNumber, err, err)
			t.Errorf("\trune = %q (%x)", test.Rune, test.Rune)
			continue
		}
		if expected, actual := test.ExpectedInt, actualInt; expected != actual {
			t.Errorf("For test #%d, expected %d, but actually got %d.", testNumber, expected, actual)
			t.Errorf("\trune = %q (%x)", test.Rune, test.Rune)
			continue
		}
		if expected, actual := test.ExpectedInt, buffer.Len(); expected != actual {
			t.Errorf("For test #%d, expected %d, but actually got %d.", testNumber, expected, actual)
			t.Errorf("\trune = %q (%x)", test.Rune, test.Rune)
			continue
		}
		for byteNumber, expected := range test.ExpectedBytes {
			if actual := buffer.Bytes()[byteNumber]; expected != actual {
				t.Errorf("For test #%d and byte #%d, expected %q (%X), but actually got %q (%X).", testNumber, byteNumber, expected, expected, actual, actual)
				t.Errorf("\trune = %q (%x)", test.Rune, test.Rune)
				continue
			}
		}
	}
}
