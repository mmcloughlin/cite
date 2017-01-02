package cite

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCitationString(t *testing.T) {
	u, err := url.Parse("https://golang.org/doc/")
	require.NoError(t, err)

	c := Citation{URL: u}
	assert.Equal(t, "https://golang.org/doc/", c.String())

	c.Extra = "extra"
	assert.Equal(t, "https://golang.org/doc/ (extra)", c.String())
}
