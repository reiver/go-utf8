package utf8

import (
	"testing"
)

func TestFormatBinary(t *testing.T) {

	tests := []struct{
		Rune     rune
		Expected string
	}{
		{
			Rune:     'A',
			Expected: "<<0b01000001>>",
		},



		{
			Rune:     'r',
			Expected: "<<0b01110010>>",
		},



		{
			Rune:     '¡',
			Expected: "<<0b11000010 ; 0b10100001>>",
		},



		{
			Rune:     '۵',
			Expected: "<<0b11011011 ; 0b10110101>>",
		},



		{
			Rune:     '‱',
			Expected: "<<0b11100010 ; 0b10000000 ; 0b10110001>>",
		},



		{
			Rune:     '≡',
			Expected: "<<0b11100010 ; 0b10001001 ; 0b10100001>>",
		},



		{
			Rune:     '🙂',
			Expected: "<<0b11110000 ; 0b10011111 ; 0b10011001 ; 0b10000010>>",
		},
	}


	for testNumber, test := range tests {

		actual := FormatBinary(test.Rune)

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, expected %q, but actually got %q.", testNumber, expected, actual)
			t.Errorf("\trune = %q (%X)", test.Rune, test.Rune)
			continue

		}
	}
}
