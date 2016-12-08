package cite

import (
	"errors"
	"net/url"
)

var ErrBadLineRange = errors.New("cite: bad line range")

type LineSelection interface {
	LineIncluded(int) bool
}

//go:generate mockery -name=LineSelection -inpkg -testonly -case=underscore

type Resource interface {
	URL() *url.URL
	Cite() Citation
	Lines() LineSelection
}

//go:generate mockery -name=Resource -inpkg -testonly -case=underscore

type ResourceBuilder func(Citation) (Resource, error)

// LineRange
type LineRange struct {
	start int
	end   int
}

func NewSingleLine(n int) LineRange {
	l, _ := NewLineRange(n, n)
	return l
}

func NewLineRange(s, e int) (LineRange, error) {
	if e < s {
		return LineRange{}, ErrBadLineRange
	}
	return LineRange{
		start: s,
		end:   e,
	}, nil
}

func (l LineRange) Start() int { return l.start }
func (l LineRange) End() int   { return l.end }

func (l LineRange) LineIncluded(n int) bool {
	return l.start <= n && n <= l.end
}
