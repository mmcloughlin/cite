package cite

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	// ErrBadLineRange is returned when bad upper and lower bounds are used to
	// construct a LineRange.
	ErrBadLineRange = errors.New("cite: bad line range")

	// ErrBadLineRangeFormat is returned for a parsing error.
	ErrBadLineRangeFormat = errors.New("cite: bad line range format")
)

// LineRange represents a range of one or more lines.
type LineRange struct {
	start int
	end   int
}

// NewSingleLine returns a LineRange object for a single line.
func NewSingleLine(n int) LineRange {
	l, _ := NewLineRange(n, n)
	return l
}

// NewLineRange returns a LineRange object for the lines between s and e
// (inclusive). It errors if the start and end do not make sense (e < s).
func NewLineRange(s, e int) (LineRange, error) {
	if e < s {
		return LineRange{}, ErrBadLineRange
	}
	return LineRange{
		start: s,
		end:   e,
	}, nil
}

// ParseLineRange constructs a LineRange from a string representation. An
// example for the typical case is "3-17" for the LineRange between 3 and 17.
// A single digit is parsed into a LineRange for a single line.
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

// Start returns the first line in the range (inclusive).
func (l LineRange) Start() int { return l.start }

// End returns the last line in the range (inclusive).
func (l LineRange) End() int { return l.end }

// LineIncluded returns whether line n is included in the range.
func (l LineRange) LineIncluded(n int) bool {
	return l.start <= n && n <= l.end
}

// NumLines returns the number of lines in the range.
func (l LineRange) NumLines() int {
	return l.end - l.start + 1
}
