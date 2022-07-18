# go-utf8

Package **utf8s** provides tools for working with Unicode encoded as UTF-8, for the Go programming language.


## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-utf8

[![GoDoc](https://godoc.org/github.com/reiver/go-utf8?status.svg)](https://godoc.org/github.com/reiver/go-utf8)


## Example

```go
var reader io.Reader

// ...

r, n, err := utf8s.ReadRune(reader)
```

```go
var writer io.Writer

// ...

var r rune

// ...

n, err := utf8s.WriteRune(w, r)
```

```go
var reader io.Reader

// ...

runeReader := utf8s.NewRuneReader(reader)

// ...

r, n, err := runeReader.ReadRune()
```

```go
var reader io.Reader

// ...

runeScanner := utf8s.NewRuneScanner(reader)

// ...

r, n, err := runeScanner.ReadRune()

// ...

err = runeScanner.UnreadRune()
```
