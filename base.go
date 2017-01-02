package cite

import (
	"fmt"
	"net/url"
)

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

type LinePredicate interface {
	LineIncluded(int) bool
}

//go:generate mockery -name=LinePredicate -inpkg -testonly -case=underscore

type LineSelection interface {
	LinePredicate
	NumLines() int
}

//go:generate mockery -name=LineSelection -inpkg -testonly -case=underscore

type Resource interface {
	URL() *url.URL
	Cite() Citation
	Lines() LineSelection
}

//go:generate mockery -name=Resource -inpkg -testonly -case=underscore

type ResourceBuilder func(Citation) (Resource, error)
