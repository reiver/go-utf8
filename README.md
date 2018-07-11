# go-utf8s

Package **utf8s** provides tools provides tools for working with Unicode encoded as UTF-8, for the Go programming language.


## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-utf8s

[![GoDoc](https://godoc.org/github.com/reiver/go-utf8s?status.svg)](https://godoc.org/github.com/reiver/go-utf8s)


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
