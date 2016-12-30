changequote(`<', `>')

# cite
Cite snippets in your godoc

[![GoDoc Reference](http://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](http://godoc.org/github.com/mmcloughlin/cite)
[![Build status](https://img.shields.io/travis/mmcloughlin/cite.svg?style=flat-square)](https://travis-ci.org/mmcloughlin/cite)
[![Coverage](https://img.shields.io/coveralls/mmcloughlin/cite.svg?style=flat-square)](https://coveralls.io/r/mmcloughlin/cite)

## Usage

To reference something, add an insert line into your `godoc`. 

```go
include(example/example.go.pre)
```

Then run

```bash
$ cite process example.go
```

It will fetch the reference and insert it into your code.

```go
include(example/example.go)
```

It should [look nice in your godoc](https://godoc.org/github.com/mmcloughlin/cite/example).
