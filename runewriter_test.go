package utf8_test

import (
	"github.com/reiver/go-utf8"

	"strings"

	"testing"
)

func TestRuneWriter(t *testing.T) {

	tests := []struct{
		Runes []rune
		Expected string
	}{
		{
			Runes: []rune{},
			Expected: "",
		},



		{
			Runes: []rune{'a','p','p','l','e'},
			Expected: "apple",
		},
		{
			Runes: []rune{'B','A','N','A','N','A'},
			Expected: "BANANA",
		},
		{
			Runes: []rune{'C','h','e','r','r','y'},
			Expected: "Cherry",
		},
		{
			Runes: []rune{'d','A','T','E'},
			Expected: "dATE",
		},



		{
			Runes: []rune{'😀','😁','😂','😃','😄','😅','😆','😇','😈','😉','😊','😋','😌','😍','😎','😏'},
			Expected: "😀😁😂😃😄😅😆😇😈😉😊😋😌😍😎😏",
		},



		{
			Runes: []rune{'H','i','!','\u2029','👾','\u2029','S','e','e',' ','y','a','\u2029'},
			Expected: "Hi!\u2029👾\u2029See ya\u2029",
		},
	}

	TestLoop: for testNumber, test := range tests {

		var buffer strings.Builder
		var total int

		runeWriter := utf8.NewRuneWriter(&buffer)

		for runeNumber, r := range test.Runes {

			n, err := runeWriter.WriteRune(r)
			if nil != err {
				t.Errorf("For test #%d and rune #%d, did not expect an error, but actually got one.", testNumber, runeNumber)
				for i, rr := range test.Runes {
					t.Logf("\t[%d] %q (%d)", i, string(rr), rr)
				}
				t.Logf("ERROR TYPE: %T", err)
				t.Logf("ERROR: %q", err)
				continue TestLoop
			}

			if expected, actual := n, utf8.RuneLength(r); expected != actual {
				t.Errorf("For test #%d and rune #%d, expected a certain number of bytes to be written, but actually wasn't.", testNumber, runeNumber)
				for i, rr := range test.Runes {
					t.Logf("\t[%d] %q (%d)", i, string(rr), rr)
				}
				t.Logf("RUNE: %q (%d)", string(r), r)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				continue TestLoop
			}

			total += n
		}

		if expected, actual := len(test.Expected), total; expected != actual {
			t.Errorf("For test #%d, expected the total number of bytes to be written to be something, but actually wasn't.", testNumber)
			for i, rr := range test.Runes {
				t.Logf("\t[%d] %q (%d)", i, string(rr), rr)
			}
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			continue TestLoop
		}

		if expected, actual := test.Expected, buffer.String(); expected != actual {
			t.Errorf("For test #%d, what what expected to be written, is not what was actually written.", testNumber)
			for i, rr := range test.Runes {
				t.Logf("\t[%d] %q (%d)", i, string(rr), rr)
			}
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue TestLoop
		}
	}
}
