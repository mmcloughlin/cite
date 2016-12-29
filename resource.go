package cite

import "net/url"

type LinePredicate interface {
	LineIncluded(int) bool
}

//go:generate mockery -name=LinePredicate -inpkg -testonly -case=underscore

type Resource interface {
	URL() *url.URL
	Cite() Citation
	Lines() LinePredicate
}

//go:generate mockery -name=Resource -inpkg -testonly -case=underscore

type ResourceBuilder func(Citation) (Resource, error)
