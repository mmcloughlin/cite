package cite

import (
	"bufio"
	"errors"
	"io"
	"net/http"
)

var ErrFetchNonStatusOK = errors.New("cite: non-200 HTTP response")

func Fetch(ref Reference) (string, error) {
	u := ref.RawFileURL()
	res, err := http.Get(u.String())
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", ErrFetchNonStatusOK
	}

	return ReadLineRange(res.Body, ref.Lines)
}

type IntRange interface {
	Start() int
	End() int
}

func ReadLineRange(r io.Reader, lines IntRange) (string, error) {
	scanner := bufio.NewScanner(r)
	output := ""
	i := 1
	for scanner.Scan() && i <= lines.End() {
		if i >= lines.Start() {
			output += scanner.Text() + "\n"
		}
		i++
	}
	return output, scanner.Err()
}
