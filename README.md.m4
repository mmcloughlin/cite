changequote(`<', `>')

# cite
Cite snippets in your godoc

[![go.dev Reference](https://img.shields.io/badge/doc-reference-007d9b?logo=go&style=flat-square)](https://pkg.go.dev/github.com/mmcloughlin/cite)
[![Build status](https://img.shields.io/travis/mmcloughlin/cite.svg?style=flat-square)](https://travis-ci.org/mmcloughlin/cite)
[![Coverage](https://img.shields.io/coveralls/mmcloughlin/cite.svg?style=flat-square)](https://coveralls.io/r/mmcloughlin/cite)
[![Go Report Card](https://goreportcard.com/badge/github.com/mmcloughlin/cite?style=flat-square)](https://goreportcard.com/report/github.com/mmcloughlin/cite)

## Install

```
go install github.com/mmcloughlin/cite/cmd/cite
```

## Usage

To reference something, add an insert line into your `godoc`. 

```go
include(example/example.go.pre)
```

Then run

```
$ cite process example.go
```

It will fetch the reference and insert it into your code.

```go
include(example/example.go)
```

It should [look nice in your godoc](https://godoc.org/github.com/mmcloughlin/cite/example).
