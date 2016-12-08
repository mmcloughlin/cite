package cite

func InsertHandler(r Resource, lines []string) ([]string, []string, error) {
	ref := Directive{
		ActionRaw: "Reference",
		Citation:  r.Cite(),
	}
	insertion := []string{ref.String()}

	snippet, err := Fetch(r)
	if err != nil {
		return nil, nil, err
	}

	insertion = append(insertion, snippet...)

	return insertion, lines[1:], nil
}
