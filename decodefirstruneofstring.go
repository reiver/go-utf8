package utf8

import (
	"io"
	"strings"
)

// DecodeFirstRuneOfString decodes the first UTF-8 encoded rune at the beginning of the string.
func DecodeFirstRuneOfString(str string) (rune, int, error) {
	var reader io.Reader = strings.NewReader(str)
	return ReadRune(reader)
}
