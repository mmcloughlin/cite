package cite

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func LineNumbersString(s, e int) string {
	out := ""
	for i := s; i <= e; i++ {
		out += strconv.Itoa(i) + "\n"
	}
	return out
}

func TestFetch(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		"https://github.com/user/repo/raw/master/LICENSE",
		httpmock.NewStringResponder(200, LineNumbersString(1, 100)),
	)

	ref := Reference{
		User:       "user",
		Repository: "repo",
		GitRef:     "master",
		Path:       "LICENSE",
		Lines:      NewLineRange(23, 40),
	}

	s, err := Fetch(ref)
	require.NoError(t, err)
	assert.Equal(t, LineNumbersString(23, 40), s)
}

// XXX http errors
// XXX odd line number cases
