package cite

import (
	"bufio"
	"errors"
	"io"
	"net/http"
)

// ErrFetchNonStatusOK is returned from Fetch on a non-200 HTTP status code.
var ErrFetchNonStatusOK = errors.New("cite: non-200 HTTP response")

// Fetch downloads a resource and extracts the referenced snippet from it.
func Fetch(r Resource) ([]string, error) {
	res, err := http.Get(r.URL().String())
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, ErrFetchNonStatusOK
	}

	return ReadLineSelection(res.Body, r.Lines())
}

// ReadLineSelection extracts the lines from r for which the given
// LinePredicate is true.
func ReadLineSelection(r io.Reader, lines LinePredicate) ([]string, error) {
	scanner := bufio.NewScanner(r)
	var output []string
	n := 1
	for scanner.Scan() {
		if lines.LineIncluded(n) {
			output = append(output, scanner.Text())
		}
		n++
	}
	return output, scanner.Err()
}
