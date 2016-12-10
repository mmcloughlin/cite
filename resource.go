package cite

import "net/url"

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
