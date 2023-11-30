package utf8_test

import (
	"testing"

	"io"
	"strings"
	"sourcecode.social/reiver/go-utf8"
)

func TestRuneScanner_Buffered(t *testing.T) {

	var s string = "🙂 apple 😈 banana 👾 cherry 🎃"

	var reader io.Reader = strings.NewReader(s)
	var runescanner utf8.RuneScanner = utf8.WrapRuneScanner(reader)

	{
		expected := 0
		actual   := runescanner.Buffered()

		if expected != actual {
			t.Errorf("The actual number of buffered-bytes is not was expected.")
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL: %d", actual)
			return
		}
	}

	{
		r, n, err := runescanner.ReadRune()
		if nil != err {
			t.Errorf("Did not expect an error when trying to .ReadRune() but actually got one.")
			t.Logf("ERROR: (%T) %s", err, err)
			return
		}

		{
			expected := '🙂'
			actual   := r

			if expected != actual {
				t.Errorf("The actual read rune is not what was expected.")
				t.Logf("EXPECTED: %q (%U)", expected, expected)
				t.Logf("ACTUAL:   %q (%U)", actual, actual)
				return
			}
		}

		{
			expected := 4
			actual   := n

			if expected != actual {
				t.Errorf("The actual number of bytes read is not what was expected.")
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				return
			}
		}
	}

	{
		expected := 0
		actual   := runescanner.Buffered()

		if expected != actual {
			t.Errorf("The actual number of buffered-bytes is not was expected.")
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL: %d", actual)
			return
		}
	}

	{
		err := runescanner.UnreadRune()
		if nil != err {
			t.Errorf("Did not expect an error when trying to .UnreadRune() but actually got one.")
			t.Logf("ERROR: (%T) %s", err, err)
			return
		}
	}

	{
		expected := 4
		actual   := runescanner.Buffered()

		if expected != actual {
			t.Errorf("The actual number of buffered-bytes is not was expected.")
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL: %d", actual)
			return
		}
	}

	{
		r, n, err := runescanner.ReadRune()
		if nil != err {
			t.Errorf("Did not expect an error when trying to .ReadRune() but actually got one.")
			t.Logf("ERROR: (%T) %s", err, err)
			return
		}

		{
			expected := '🙂'
			actual   := r

			if expected != actual {
				t.Errorf("The actual read rune is not what was expected.")
				t.Logf("EXPECTED: %q (%U)", expected, expected)
				t.Logf("ACTUAL:   %q (%U)", actual, actual)
				return
			}
		}

		{
			expected := 4
			actual   := n

			if expected != actual {
				t.Errorf("The actual number of bytes read is not what was expected.")
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				return
			}
		}
	}

	{
		expected := 0
		actual   := runescanner.Buffered()

		if expected != actual {
			t.Errorf("The actual number of buffered-bytes is not was expected.")
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL: %d", actual)
			return
		}
	}

	{
		r, n, err := runescanner.ReadRune()
		if nil != err {
			t.Errorf("Did not expect an error when trying to .ReadRune() but actually got one.")
			t.Logf("ERROR: (%T) %s", err, err)
			return
		}

		{
			expected := ' '
			actual   := r

			if expected != actual {
				t.Errorf("The actual read rune is not what was expected.")
				t.Logf("EXPECTED: %q (%U)", expected, expected)
				t.Logf("ACTUAL:   %q (%U)", actual, actual)
				return
			}
		}

		{
			expected := 1
			actual   := n

			if expected != actual {
				t.Errorf("The actual number of bytes read is not what was expected.")
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				return
			}
		}
	}

	{
		expected := 0
		actual   := runescanner.Buffered()

		if expected != actual {
			t.Errorf("The actual number of buffered-bytes is not was expected.")
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL: %d", actual)
			return
		}
	}

	{
		err := runescanner.UnreadRune()
		if nil != err {
			t.Errorf("Did not expect an error when trying to .UnreadRune() but actually got one.")
			t.Logf("ERROR: (%T) %s", err, err)
			return
		}
	}

	{
		expected := 1
		actual   := runescanner.Buffered()

		if expected != actual {
			t.Errorf("The actual number of buffered-bytes is not was expected.")
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL: %d", actual)
			return
		}
	}
}
