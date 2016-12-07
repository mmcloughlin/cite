package cite

import (
	"bufio"
	"errors"
	"io"
	"net/http"
)

var ErrFetchNonStatusOK = errors.New("cite: non-200 HTTP response")

func Fetch(r Resource) (string, error) {
	res, err := http.Get(r.URL().String())
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", ErrFetchNonStatusOK
	}

	return ReadLineSelection(res.Body, r.Lines())
}

func ReadLineSelection(r io.Reader, lines LineSelection) (string, error) {
	scanner := bufio.NewScanner(r)
	output := ""
	n := 1
	for scanner.Scan() {
		if lines.LineIncluded(n) {
			output += scanner.Text() + "\n"
		}
		n++
	}
	return output, scanner.Err()
}
