package cite

import (
	"bufio"
	"errors"
	"io"
	"net/http"
)

var ErrFetchNonStatusOK = errors.New("cite: non-200 HTTP response")

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

func ReadLineSelection(r io.Reader, lines LineSelection) ([]string, error) {
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
