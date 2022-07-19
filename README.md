# go-utf8

Package **utf8** provides tools for working with Unicode encoded as UTF-8, for the Go programming language.


## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-utf8

[![GoDoc](https://godoc.org/github.com/reiver/go-utf8?status.svg)](https://godoc.org/github.com/reiver/go-utf8)


## Example

```go
var reader io.Reader

// ...

r, n, err := utf8.ReadRune(reader)
```

```go
var writer io.Writer

// ...

var r rune

// ...

n, err := utf8.WriteRune(w, r)
```

```go
var reader io.Reader

// ...

runeReader := utf8.RuneReaderWrap(reader)

// ...

r, n, err := runeReader.ReadRune()
```

```go
var reader io.Reader

// ...

runeScanner := utf8.RuneScannerWrap(reader)

// ...

r, n, err := runeScanner.ReadRune()

// ...

err = runeScanner.UnreadRune()
```
