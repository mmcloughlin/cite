package cite

func FormatSnippet(lines []string) []string {
	output := make([]string, len(lines)+2)
	for i, line := range lines {
		output[i+1] = "\t" + line
	}
	return output
}

func InsertHandler(r Resource, lines []string) ([]string, []string, error) {
	ref := Directive{
		ActionRaw: "Reference",
		Citation:  r.Cite(),
	}
	insertion := []string{" " + ref.String()}

	snippet, err := Fetch(r)
	if err != nil {
		return nil, nil, err
	}

	formatted := FormatSnippet(snippet)
	insertion = append(insertion, formatted...)

	return insertion, lines[1:], nil
}
