package cite

import "net/url"

type PlainResource struct {
	FileURL   *url.URL
	LineRange LineRange
}

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

func (p PlainResource) URL() *url.URL {
	return p.FileURL
}

func (p PlainResource) Cite() Citation {
	return Citation{
		URL:   p.FileURL,
		Extra: p.LineRange.String(),
	}
}

func (p PlainResource) Lines() LineSelection {
	return p.LineRange
}
