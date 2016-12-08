package cite

import (
	"net/http"
	"net/url"
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

	u, err := url.Parse("http://idk.com/doc.txt")
	require.NoError(t, err)

	lines, err := NewLineRange(3, 5)
	require.NoError(t, err)

	r := &MockResource{}
	r.On("URL").Return(u).Once()
	r.On("Lines").Return(lines).Once()

	httpmock.RegisterResponder(
		http.MethodGet,
		u.String(),
		httpmock.NewStringResponder(200, LineNumbersString(1, 10)),
	)

	s, err := Fetch(r)

	require.NoError(t, err)
	assert.Equal(t, []string{"3", "4", "5"}, s)
}

// XXX http errors
// XXX odd line number cases
