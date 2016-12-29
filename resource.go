package cite

import "net/url"

type LinePredicate interface {
	LineIncluded(int) bool
}

//go:generate mockery -name=LinePredicate -inpkg -testonly -case=underscore

type LineSelection interface {
	LinePredicate
	NumLines() int
}

type Resource interface {
	URL() *url.URL
	Cite() Citation
	Lines() LineSelection
}

//go:generate mockery -name=Resource -inpkg -testonly -case=underscore

type ResourceBuilder func(Citation) (Resource, error)
