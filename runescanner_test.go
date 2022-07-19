package utf8

import (
	"io"
	"strings"

	"testing"
)

func TestRuneScanner(t *testing.T) {

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
		runeReader := RuneScannerWrap(test.Reader)

		actualRune, actualInt, err := runeReader.ReadRune()
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

func TestRuneScanners(t *testing.T) {

	tests := []struct{
		Reader     io.Reader
		Expected []rune
	}{
		{ // 0
			Reader: strings.NewReader("a"),
			Expected: []rune{'a'},
		},
		{ // 1
			Reader: strings.NewReader("ap"),
			Expected: []rune{'a','p'},
		},
		{ // 2
			Reader: strings.NewReader("app"),
			Expected: []rune{'a','p','p'},
		},
		{ // 3
			Reader: strings.NewReader("appl"),
			Expected: []rune{'a','p','p','l'},
		},
		{ // 4
			Reader: strings.NewReader("apple"),
			Expected: []rune{'a','p','p','l','e'},
		},



		{ // 5
			Reader: strings.NewReader("b"),
			Expected: []rune{'b'},
		},
		{ // 6
			Reader: strings.NewReader("ba"),
			Expected: []rune{'b','a'},
		},
		{ // 7
			Reader: strings.NewReader("ban"),
			Expected: []rune{'b','a','n'},
		},
		{ // 8
			Reader: strings.NewReader("bana"),
			Expected: []rune{'b','a','n','a'},
		},
		{ // 9
			Reader: strings.NewReader("banan"),
			Expected: []rune{'b','a','n','a','n'},
		},
		{ // 10
			Reader: strings.NewReader("banana"),
			Expected: []rune{'b','a','n','a','n','a'},
		},



		{ // 11
			Reader: strings.NewReader("c"),
			Expected: []rune{'c'},
		},
		{ // 12
			Reader: strings.NewReader("ch"),
			Expected: []rune{'c','h'},
		},
		{ // 13
			Reader: strings.NewReader("che"),
			Expected: []rune{'c','h','e'},
		},
		{ // 14
			Reader: strings.NewReader("cher"),
			Expected: []rune{'c','h','e','r'},
		},
		{ // 15
			Reader: strings.NewReader("cherr"),
			Expected: []rune{'c','h','e','r','r'},
		},
		{ // 16
			Reader: strings.NewReader("cherry"),
			Expected: []rune{'c','h','e','r','r','y'},
		},



		{ // 17
			Reader: strings.NewReader("A"),
			Expected: []rune{'A'},
		},



		{ // 18
			Reader: strings.NewReader("r"),
			Expected: []rune{'r'},
		},



		{ // 19
			Reader: strings.NewReader("¡"),
			Expected: []rune{'¡'},
		},
		{ // 20
			Reader: strings.NewReader("¡!"),
			Expected: []rune{'¡','!'},
		},



		{ // 21
			Reader: strings.NewReader("۵"),
			Expected: []rune{'۵'},
		},
		{ // 22
			Reader: strings.NewReader("۵5"),
			Expected: []rune{'۵','5'},
		},



		{ // 23
			Reader: strings.NewReader("‱"),
			Expected: []rune{'‱'},
		},
		{ // 24
			Reader: strings.NewReader("‱%"),
			Expected: []rune{'‱','%'},
		},



		{ // 25
			Reader: strings.NewReader("≡"),
			Expected: []rune{'≡'},
		},
		{ // 26
			Reader: strings.NewReader("≡="),
			Expected: []rune{'≡', '='},
		},



		{ // 27
			Reader: strings.NewReader("𐏕"),
			Expected: []rune{'𐏕'},
		},
		{ // 28
			Reader: strings.NewReader("𐏕100"),
			Expected: []rune{'𐏕','1','0','0'},
		},



		{ // 29
			Reader: strings.NewReader("🙂"),
			Expected: []rune{'🙂'},
		},
		{ // 30
			Reader: strings.NewReader("🙂:-)"),
			Expected: []rune{'🙂',':','-',')'},
		},



		{ // 31
			Reader: strings.NewReader("\u0000"),
			Expected: []rune{0x0},
		},
		{ // 32
			Reader: strings.NewReader("\u0001"),
			Expected: []rune{0x1},
		},
		{ // 33
			Reader: strings.NewReader("\u007e"),
			Expected: []rune{0x7e},
		},
		{ // 34
			Reader: strings.NewReader("\u007f"),
			Expected: []rune{0x7f},
		},



		{ // 35
			Reader: strings.NewReader("\u0080"),
			Expected: []rune{0x80},
		},
		{ // 36
			Reader: strings.NewReader("\u0081"),
			Expected: []rune{0x81},
		},
		{ // 37
			Reader: strings.NewReader("\u07fe"),
			Expected: []rune{0x7fe},
		},
		{ // 38
			Reader: strings.NewReader("\u07ff"),
			Expected: []rune{0x7ff},
		},



		{ // 39
			Reader: strings.NewReader("\u0800"),
			Expected: []rune{0x800},
		},
		{ // 40
			Reader: strings.NewReader("\u0801"),
			Expected: []rune{0x801},
		},
		{ // 41
			Reader: strings.NewReader("\ufffe"),
			Expected: []rune{0xfffe},
		},
		{ // 42
			Reader: strings.NewReader("\uffff"),
			Expected: []rune{0xffff},
		},



		{ // 43
			Reader: strings.NewReader("\U00010000"),
			Expected: []rune{0x10000},
		},
		{ // 44
			Reader: strings.NewReader("\U00010001"),
			Expected: []rune{0x10001},
		},
		{ // 45
			Reader: strings.NewReader("\U0010fffe"),
			Expected: []rune{0x10fffe},
		},
		{ // 46
			Reader: strings.NewReader("\U0010ffff"),
			Expected: []rune{0x10ffff},
		},
	}


	TestLoop: for testNumber, test := range tests {

		var runeNumber int
		for {
			runeReader := RuneScannerWrap(test.Reader)

			actualRune, actualInt, err := runeReader.ReadRune()
			if nil != err && io.EOF != err {
				t.Errorf("For test #%d and rune #%d, did not expect an error, but actually got one: (%T) %q", testNumber, runeNumber, err, err)
				t.Errorf("\tEXPECTED: %s", FormatBinary(test.Expected[runeNumber]))
				t.Errorf("\tACTUAL:   %s", FormatBinary(actualRune))
				continue TestLoop
			}
			if io.EOF == err {
				if expected, actual := len(test.Expected), runeNumber; expected != actual {
					t.Errorf("For test #%d and rune #%d, expected %d, but actually got %d.", testNumber, runeNumber, expected, actual)
				}

				break
			}
			if expected, actual := test.Expected[runeNumber], actualRune; expected != actual {
				t.Errorf("For test #%d and rune #%d, expected %q (0x%X), but actually got %q (0x%X).", testNumber, runeNumber, expected, expected, actual, actual)
				t.Errorf("\tEXPECTED: %s", FormatBinary(test.Expected[runeNumber]))
				t.Errorf("\tACTUAL:   %s", FormatBinary(actualRune))
				continue TestLoop
			}
			if expected, actual := RuneLength(test.Expected[runeNumber]), actualInt; expected != actual {
				t.Errorf("For test #%d and rune #%d, expected %d, but actually got %d.", testNumber, runeNumber, expected, actual)
				t.Errorf("\tEXPECTED: %s", FormatBinary(test.Expected[runeNumber]))
				t.Errorf("\tACTUAL:   %s", FormatBinary(actualRune))
				continue TestLoop
			}

			runeNumber++
		}
	}
}

func TestRuneScannerUnread(t *testing.T) {

	tests := []struct{
		Reader io.Reader
		Instructions []rune
		ExpectedRune []rune
		ExpectedSize []int
	}{
		{
			Reader: strings.NewReader("a ≡ b\r\n۰۱۲۳۴۵۶۷۸۹ \U00010001"),
			Instructions: []rune{'r', 'u', 'r', 'u', 'u', 'r', 'r', 'u', 'r', 'r', 'u', 'r', 'u', 'u', 'r', 'r', 'r',  'r',  'r', 'r', 'r', 'r', 'r', 'r', 'r', 'u', 'u', 'u', 'r', 'r', 'r', 'r', 'r', 'r', 'r'},
			ExpectedRune: []rune{'a',      'a',           'a', ' ',      ' ', '≡',      '≡',           '≡', ' ', 'b', '\r', '\n', '۰', '۱', '۲', '۳', '۴', '۵',                '۵', '۶', '۷', '۸', '۹', ' ', '\U00010001'},
			ExpectedSize: []int{  1,        1,             1,   1,        1,   3,        3,             3,   1,   1,    1,    1,   2,   2,   2,   2,   2,   2,                  2,   2,   2,   2,   2,   1,   4},
		},
	}


	TestLoop: for testNumber, test := range tests {

		runeScanner := RuneScannerWrap(test.Reader)

		var readCount int
		for instructionNumber, instruction := range test.Instructions {

			switch instruction {
			case 'r': // =   rea

				actualRune, actualSize, err := runeScanner.ReadRune()
				if nil != err {
					t.Errorf("For test #%d and instruction #%d, did not expected an error, but actually got one: (%T) %q", testNumber, instructionNumber, err, err)
					continue TestLoop
				}

				expectedRune := test.ExpectedRune[readCount]
				expectedSize := test.ExpectedSize[readCount]

				if expected, actual := expectedRune, actualRune; expected != actual {
					t.Errorf("For test #%d and instruction #%d, expected rune %q / %d, but actually got %q / %d", testNumber, instructionNumber, expected, expected, actual, actual)
					continue TestLoop
				}

				if expected, actual := expectedSize, actualSize; expected != actual {
					t.Errorf("For test #%d and instruction #%d, for rune %q / %d expected size %d, but actually got size %d", testNumber, instructionNumber, expectedRune, expectedRune, expected, actual)
					continue TestLoop
				}

				readCount++
			case 'u': // = unread
				if err := runeScanner.UnreadRune(); nil != err {
					t.Errorf("For test #%d and instruction #%d, did not expected an error, but actually got one: (%T) %q", testNumber, instructionNumber, err, err)
					continue TestLoop
				}
			default:
				t.Errorf("For test #%d and instruction #%d, UNKNOWN INSTRUCTION!!!:... %q", testNumber, instructionNumber, instruction)
				continue TestLoop
			}

		}

	}
}
