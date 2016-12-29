package cite

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	ErrBadLineRange       = errors.New("cite: bad line range")
	ErrBadLineRangeFormat = errors.New("cite: bad line range format")
)

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

func ParseLineRange(s string) (LineRange, error) {
	parts := strings.Split(s, "-")
	if len(parts) > 2 {
		return LineRange{}, ErrBadLineRangeFormat
	}

	start, err := strconv.Atoi(parts[0])
	if err != nil {
		return LineRange{}, err
	}

	if len(parts) == 1 {
		return NewSingleLine(start), nil
	}

	end, err := strconv.Atoi(parts[1])
	if err != nil {
		return LineRange{}, err
	}

	return NewLineRange(start, end)
}

func (l LineRange) String() string {
	if l.start == l.end {
		return strconv.Itoa(l.start)
	}
	return fmt.Sprintf("%d-%d", l.start, l.end)
}

func (l LineRange) Start() int { return l.start }
func (l LineRange) End() int   { return l.end }

func (l LineRange) LineIncluded(n int) bool {
	return l.start <= n && n <= l.end
}

func (l LineRange) NumLines() int {
	return l.end - l.start + 1
}
