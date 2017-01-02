package cite

import "errors"

var ErrShortReference = errors.New("cite: existing reference too short")

func FormatSnippet(lines []string) []string {
	output := make([]string, len(lines)+2)
	for i, line := range lines {
		output[i+1] = "\t" + line
	}
	return output
}

func InsertHandler(r Resource, lines []string) ([]string, []string, error) {
	snippet, err := Fetch(r)
	if err != nil {
		return nil, nil, err
	}

	ref := Directive{
		ActionRaw: "Reference",
		Citation:  r.Cite(),
	}
	insertion := []string{" " + ref.String()}

	formatted := FormatSnippet(snippet)
	insertion = append(insertion, formatted...)

	return insertion, lines[1:], nil
}

func SkipReferenceHandler(r Resource, lines []string) ([]string, []string, error) {
	snippetLength := r.Lines().NumLines()
	referenceLength := snippetLength + 3
	if len(lines) < referenceLength {
		return nil, nil, ErrShortReference
	}
	return lines[:referenceLength], lines[referenceLength:], nil
}
