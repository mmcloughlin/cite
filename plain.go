package cite

import "net/url"

// PlainResource is a snippet from an arbitrary URL.
type PlainResource struct {
	FileURL   *url.URL
	LineRange LineRange
}

// BuildPlainResourceFromCitation builds a PlainResource from a reference to
// it. The line range is represented in the Extra field of the Citation.
func BuildPlainResourceFromCitation(c Citation) (Resource, error) {
	lines, err := ParseLineRange(c.Extra)
	if err != nil {
		return nil, err
	}
	return PlainResource{
		FileURL:   c.URL,
		LineRange: lines,
	}, nil
}

// URL returns the URL of the whole file.
func (p PlainResource) URL() *url.URL {
	return p.FileURL
}

// Cite returns a reference to the resource.
func (p PlainResource) Cite() Citation {
	return Citation{
		URL:   p.FileURL,
		Extra: p.LineRange.String(),
	}
}

// Lines returns the selection of lines in the snippet.
func (p PlainResource) Lines() LineSelection {
	return p.LineRange
}
