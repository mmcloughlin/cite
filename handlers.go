package cite

import "errors"

// ErrShortReference occurs when the snippet following a reference is too
// short. That is, it does not contain as many lines as the reference
// specifies.
var ErrShortReference = errors.New("cite: existing reference too short")

// FormatSnippet formats the raw lines into the lines to be inserted into the
// comment. This includes indendation and blank lines above and below.
func FormatSnippet(lines []string) []string {
	output := make([]string, len(lines)+2)
	for i, line := range lines {
		output[i+1] = "\t" + line
	}
	return output
}

// InsertHandler handles "insert" actions by fetching and inserting the
// snippet into the comment below.
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

// SkipReferenceHandler handles "reference" actions by calculating the length
// of the snippet and skipping that many lines below.
func SkipReferenceHandler(r Resource, lines []string) ([]string, []string, error) {
	snippetLength := r.Lines().NumLines()
	referenceLength := snippetLength + 3
	if len(lines) < referenceLength {
		return nil, nil, ErrShortReference
	}
	return lines[:referenceLength], lines[referenceLength:], nil
}
