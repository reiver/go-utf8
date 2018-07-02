package utf8s

import (
	"io"
	"strings"

	"testing"
)

func TestReadRune(t *testing.T) {

	tests := []struct{
		Reader       io.Reader
		ExpectedRune rune
		ExpectedInt  int
	}{
		{
			Reader: strings.NewReader("a"),
			ExpectedRune: 'a',
			ExpectedInt:   1,
		},
		{
			Reader: strings.NewReader("ap"),
			ExpectedRune: 'a',
			ExpectedInt:   1,
		},
		{
			Reader: strings.NewReader("app"),
			ExpectedRune: 'a',
			ExpectedInt:   1,
		},
		{
			Reader: strings.NewReader("appl"),
			ExpectedRune: 'a',
			ExpectedInt:   1,
		},
		{
			Reader: strings.NewReader("apple"),
			ExpectedRune: 'a',
			ExpectedInt:   1,
		},



		{
			Reader: strings.NewReader("b"),
			ExpectedRune: 'b',
			ExpectedInt:   1,
		},
		{
			Reader: strings.NewReader("ba"),
			ExpectedRune: 'b',
			ExpectedInt:   1,
		},
		{
			Reader: strings.NewReader("ban"),
			ExpectedRune: 'b',
			ExpectedInt:   1,
		},
		{
			Reader: strings.NewReader("bana"),
			ExpectedRune: 'b',
			ExpectedInt:   1,
		},
		{
			Reader: strings.NewReader("banan"),
			ExpectedRune: 'b',
			ExpectedInt:   1,
		},
		{
			Reader: strings.NewReader("banana"),
			ExpectedRune: 'b',
			ExpectedInt:   1,
		},



		{
			Reader: strings.NewReader("c"),
			ExpectedRune: 'c',
			ExpectedInt:   1,
		},
		{
			Reader: strings.NewReader("ch"),
			ExpectedRune: 'c',
			ExpectedInt:   1,
		},
		{
			Reader: strings.NewReader("che"),
			ExpectedRune: 'c',
			ExpectedInt:   1,
		},
		{
			Reader: strings.NewReader("cher"),
			ExpectedRune: 'c',
			ExpectedInt:   1,
		},
		{
			Reader: strings.NewReader("cherr"),
			ExpectedRune: 'c',
			ExpectedInt:   1,
		},
		{
			Reader: strings.NewReader("cherry"),
			ExpectedRune: 'c',
			ExpectedInt:   1,
		},



		{
			Reader: strings.NewReader("A"),
			ExpectedRune: 'A',
			ExpectedInt:   1,
		},



		{
			Reader: strings.NewReader("r"),
			ExpectedRune: 'r',
			ExpectedInt:   1,
		},



		{
			Reader: strings.NewReader("¡"),
			ExpectedRune: '¡',
			ExpectedInt:   2,
		},
		{
			Reader: strings.NewReader("¡!"),
			ExpectedRune: '¡',
			ExpectedInt:   2,
		},



		{
			Reader: strings.NewReader("۵"),
			ExpectedRune: '۵',
			ExpectedInt:   2,
		},
		{
			Reader: strings.NewReader("۵5"),
			ExpectedRune: '۵',
			ExpectedInt:   2,
		},



		{
			Reader: strings.NewReader("‱"),
			ExpectedRune: '‱',
			ExpectedInt:   3,
		},
		{
			Reader: strings.NewReader("‱%"),
			ExpectedRune: '‱',
			ExpectedInt:   3,
		},



		{
			Reader: strings.NewReader("≡"),
			ExpectedRune: '≡',
			ExpectedInt:   3,
		},
		{
			Reader: strings.NewReader("≡="),
			ExpectedRune: '≡',
			ExpectedInt:   3,
		},



		{
			Reader: strings.NewReader("𐏕"),
			ExpectedRune: '𐏕',
			ExpectedInt:   4,
		},
		{
			Reader: strings.NewReader("𐏕100"),
			ExpectedRune: '𐏕',
			ExpectedInt:   4,
		},



		{
			Reader: strings.NewReader("🙂"),
			ExpectedRune: '🙂',
			ExpectedInt:   4,
		},
		{
			Reader: strings.NewReader("🙂:-)"),
			ExpectedRune: '🙂',
			ExpectedInt:   4,
		},



		{
			Reader: strings.NewReader("\u0000"),
			ExpectedRune:  0x0,
			ExpectedInt:   1,
		},
		{
			Reader: strings.NewReader("\u0001"),
			ExpectedRune:  0x1,
			ExpectedInt:   1,
		},
		{
			Reader: strings.NewReader("\u007e"),
			ExpectedRune:  0x7e,
			ExpectedInt:   1,
		},
		{
			Reader: strings.NewReader("\u007f"),
			ExpectedRune:  0x7f,
			ExpectedInt:   1,
		},



		{
			Reader: strings.NewReader("\u0080"),
			ExpectedRune:  0x80,
			ExpectedInt:   2,
		},
		{
			Reader: strings.NewReader("\u0081"),
			ExpectedRune:  0x81,
			ExpectedInt:   2,
		},
		{
			Reader: strings.NewReader("\u07fe"),
			ExpectedRune:  0x7fe,
			ExpectedInt:   2,
		},
		{
			Reader: strings.NewReader("\u07ff"),
			ExpectedRune:  0x7ff,
			ExpectedInt:   2,
		},



		{
			Reader: strings.NewReader("\u0800"),
			ExpectedRune:  0x800,
			ExpectedInt:   3,
		},
		{
			Reader: strings.NewReader("\u0801"),
			ExpectedRune:  0x801,
			ExpectedInt:   3,
		},
		{
			Reader: strings.NewReader("\ufffe"),
			ExpectedRune:  0xfffe,
			ExpectedInt:   3,
		},
		{
			Reader: strings.NewReader("\uffff"),
			ExpectedRune:  0xffff,
			ExpectedInt:   3,
		},



		{
			Reader: strings.NewReader("\U00010000"),
			ExpectedRune:  0x10000,
			ExpectedInt:   4,
		},
		{
			Reader: strings.NewReader("\U00010001"),
			ExpectedRune:  0x10001,
			ExpectedInt:   4,
		},
		{
			Reader: strings.NewReader("\U0010fffe"),
			ExpectedRune:  0x10fffe,
			ExpectedInt:   4,
		},
		{
			Reader: strings.NewReader("\U0010ffff"),
			ExpectedRune:  0x10ffff,
			ExpectedInt:   4,
		},
	}


	for testNumber, test := range tests {

		actualRune, actualInt, err := ReadRune(test.Reader)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q", testNumber, err, err)
			t.Errorf("\tEXPECTED: %s", FormatBinary(test.ExpectedRune))
			t.Errorf("\tACTUAL:   %s", FormatBinary(actualRune))
			continue
		}
		if expected, actual := test.ExpectedRune, actualRune; expected != actual {
			t.Errorf("For test #%d, expected %q (0x%X), but actually got %q (0x%X).", testNumber, expected, expected, actual, actual)
			t.Errorf("\tEXPECTED: %s", FormatBinary(test.ExpectedRune))
			t.Errorf("\tACTUAL:   %s", FormatBinary(actualRune))
			continue
		}
		if expected, actual := test.ExpectedInt, actualInt; expected != actual {
			t.Errorf("For test #%d, expected %d, but actually got %d.", testNumber, expected, actual)
			t.Errorf("\tEXPECTED: %s", FormatBinary(test.ExpectedRune))
			t.Errorf("\tACTUAL:   %s", FormatBinary(actualRune))
			continue
		}

	}
}
