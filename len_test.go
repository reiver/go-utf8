package utf8s

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
	}


	for testNumber, test := range tests {

		actual := Len(test.Datum)
		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, expected %d, but actually got %d.", testNumber, expected, actual)
			continue
		}
	}
}
