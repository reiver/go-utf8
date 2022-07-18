package utf8

import (
	"testing"
)

func TestLen(t *testing.T) {

	tests := []struct{
		Datum rune
		Expected int
	}{
		{
			Datum:   'A',
			Expected: 1,
		},
		{
			Datum:   'r',
			Expected: 1,
		},
		{
			Datum:   '¡',
			Expected: 2,
		},
		{
			Datum:   '۵',
			Expected: 2,
		},
		{
			Datum:   '‱',
			Expected: 3,
		},
		{
			Datum:   '≡',
			Expected: 3,
		},
		{
			Datum:   '𐏕',
			Expected: 4,
		},
		{
			Datum:   '🙂',
			Expected: 4,
		},



		{
			Datum:   '\u0000',
			Expected: 1,
		},
		{
			Datum:   '\u0001',
			Expected: 1,
		},
		{
			Datum:   '\u007e',
			Expected: 1,
		},
		{
			Datum:   '\u007f',
			Expected: 1,
		},



		{
			Datum:   '\u0080',
			Expected: 2,
		},
		{
			Datum:   '\u0081',
			Expected: 2,
		},
		{
			Datum:   '\u07fe',
			Expected: 2,
		},
		{
			Datum:   '\u07ff',
			Expected: 2,
		},



		{
			Datum:   '\u0800',
			Expected: 3,
		},
		{
			Datum:   '\u0801',
			Expected: 3,
		},
		{
			Datum:   '\ufffe',
			Expected: 3,
		},
		{
			Datum:   '\uffff',
			Expected: 3,
		},



		{
			Datum:   '\U00010000',
			Expected: 4,
		},
		{
			Datum:   '\U00010001',
			Expected: 4,
		},
		{
			Datum:   '\U0010fffe',
			Expected: 4,
		},
		{
			Datum:   '\U0010ffff',
			Expected: 4,
		},
	}


	for testNumber, test := range tests {

		actual := Len(test.Datum)
		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, expected %d, but actually got %d.", testNumber, expected, actual)
			continue
		}
	}
}
