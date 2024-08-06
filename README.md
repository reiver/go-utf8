# go-utf8

Package **utf8** implements encoding and decoding of UTF-8, for the Go programming language.

This package is meant to be a replacement for Go's built-in `"unicode/utf8"` package.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-utf8

[![GoDoc](https://godoc.org/github.com/reiver/go-utf8?status.svg)](https://godoc.org/github.com/reiver/go-utf8)

## Reading a Single UTF-8 Character

This is the simplest way of reading a single UTF-8 character.

```go
var reader io.Reader

// ...

r, n, err := utf8.ReadRune(reader)
```
## Write a Single UTF-8 Character

This is the simplest way of writing a single UTF-8 character.

```go
var writer io.Writer

// ...

var r rune

// ...

n, err := utf8.WriteRune(w, r)
```
## io.RuneReader

This is how you can create an `io.RuneReader`:

```go
var reader io.Reader

// ...

var runeReader io.RuneReader = utf8.NewRuneReader(reader)

// ...

r, n, err := runeReader.ReadRune()
```
## io.RuneScanner

This is how you can create an `io.RuneScanner`:

```go
var reader io.Reader

// ...

var runeScanner io.RuneScanner := utf8.NewRuneScanner(reader)

// ...

r, n, err := runeScanner.ReadRune()

// ...

err = runeScanner.UnreadRune()
```

## UTF-8

UTF-8 is a variable length encoding of Unicode.
An encoding of a single Unicode code point can be from 1 to 4 bytes longs.

Some examples of UTF-8 encoding of Unicode code points are:

<table>
	<thead>
		<tr>
			<td colspan="4" align="center">UTF-8 encoding</td>
			<td rowspan="2">value</td>
			<td rowspan="2">code point</td>
			<td rowspan="2">decimal</td>
			<td rowspan="2">binary</td>
			<td rowspan="2">name</td>
		</tr>
		<tr>
			<td>byte 1</td>
			<td>byte 2</td>
			<td>byte 3</td>
			<td>byte 4</td>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td><code>0b0,1000001</code></td>
			<td></td>
			<td></td>
			<td></td>
			<td>A</td>
			<td>U+0041</td>
			<td>65</td>
			<td><code>0b0000,0000,0100,0001</code></td>
			<td>LATIN CAPITAL LETTER A</td>
		</tr>
		<tr>
			<td><code>0b0,1110010</code></td>
			<td></td>
			<td></td>
			<td></td>
			<td>r</td>
			<td>U+0072</td>
			<td>114</td>
			<td><code>0b0000,0000,0111,0010</code></td>
			<td>LATIN SMALL LETTER R</td>
		</tr>
		<tr>
			<td><code>0b110,00010</code></td>
			<td><code>0b10,100001</code></td>
			<td></td>
			<td></td>
			<td>¡</td>
			<td>U+00A1</td>
			<td>161</td>
			<td><code>0b0000,0000,1010,0001</code></td>
			<td>INVERTED EXCLAMATION MARK</td>
		</tr>
		<tr>
			<td><code>0b110,11011</code></td>
			<td><code>0b10,110101</code></td>
			<td></td>
			<td></td>
			<td>۵</td>
			<td>U+06F5</td>
			<td>1781</td>
			<td><code>0b0000,0110,1111,0101</code></td>
			<td>EXTENDED ARABIC-INDIC DIGIT FIVE</td>
		</tr>
		<tr>
			<td><code>0b1110,0010</code></td>
			<td><code>0b10,000000</code></td>
			<td><code>0b10,110001</code></td>
			<td></td>
			<td>‱</td>
			<td>U+2031</td>
			<td>8241</td>
			<td><code>0b0010,0000,0011,0001</code></td>
			<td>PER TEN THOUSAND SIGN</td>
		</tr>
		<tr>
			<td><code>0b1110,0010</code></td>
			<td><code>0b10,001001</code></td>
			<td><code>0b10,100001</code></td>
			<td></td>
			<td>≡</td>
			<td>U+2261</td>
			<td>8801</td>
			<td><code>0b0010,0010,0110,0001</code></td>
			<td>IDENTICAL TO</td>
		</tr>
		<tr>
			<td><code>0b11110,000</code></td>
			<td><code>0b10,010000</code></td>
			<td><code>0b10,001111</code></td>
			<td><code>0b10,010101</code></td>
			<td>𐏕</td>
			<td>U+000103D5</td>
			<td>66517</td>
			<td><code>b0001,0000,0011,1101,0101</code></td>
			<td>OLD PERSIAN NUMBER HUNDRED</td>
		</tr>
		<tr>
			<td><code>0b11110,000</code></td>
			<td><code>0b10,011111</code></td>
			<td><code>0b10,011001</code></td>
			<td><code>0b10,000010</code></td>
			<td>🙂</td>
			<td>U+0001F642</td>
			<td>128578</td>
			<td><code>0b0001,1111,0110,0100,0010</code></td>
			<td>SLIGHTLY SMILING FACE</td>
		</tr>
	</tbody>
</table>

## UTF-8 Versus ASCII

UTF-8 was (partially) designed to be backwards compatible with 7-bit ASCII.

Thus, all 7-bit ASCII is valid UTF-8.

## UTF-8 Encoding

Since, at least as of 2003, Unicode fits into 21 bits, and thus UTF-8 was designed to support at most 21 bits of information.

This is done as described in the following table:

| # of bytes | # bits for code point | 1st code point |  last code point |   byte 1   |   byte 2   |   byte 3   |   byte 4   |
|------------|-----------------------|----------------|------------------|------------|------------|------------|------------|
|     1      |            7          |    U+000000    |     U+00007F     | `0xxxxxxx` |            |            |            |
|     2      |           11          |    U+000080    |     U+0007FF     | `110xxxxx` | `10xxxxxx` |            |            |
|     3      |           16          |    U+000800    |     U+00FFFF     | `1110xxxx` | `10xxxxxx` | `10xxxxxx` |            |
|     4      |           21          |    U+010000    |     U+10FFFF     | `11110xxx` | `10xxxxxx` | `10xxxxxx` | `10xxxxxx` |
```
