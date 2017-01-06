package cite

import (
	"fmt"
	"net/url"
)

// Citation is a reference to an external resource, perhaps specifiying
// a particular segment of it with some extra data.
type Citation struct {
	URL   *url.URL
	Extra string
}

func (c Citation) String() string {
	if c.Extra == "" {
		return c.URL.String()
	}
	return fmt.Sprintf("%v (%s)", c.URL, c.Extra)
}

// LinePredicate has a boolean function on line numbers, specifying some
// subset of interest.
type LinePredicate interface {
	LineIncluded(int) bool
}

//go:generate mockery -name=LinePredicate -inpkg -testonly -case=underscore

// LineSelection specifies a subset of lines of interest (like LinePredicate)
// of a known size.
type LineSelection interface {
	LinePredicate
	NumLines() int
}

//go:generate mockery -name=LineSelection -inpkg -testonly -case=underscore

// Resource is (a subset of) an external resource which can be cited.
type Resource interface {
	URL() *url.URL
	Cite() Citation
	Lines() LineSelection
}

//go:generate mockery -name=Resource -inpkg -testonly -case=underscore

// ResourceBuilder is a function that constructs a resource from a reference
// to it.
type ResourceBuilder func(Citation) (Resource, error)
