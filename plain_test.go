package cite

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBuildPlainResourceFromCitation(t *testing.T) {
	u, err := url.Parse("http://plain.com/path")
	require.NoError(t, err)
	citation := Citation{
		URL:   u,
		Extra: "3-7",
	}
	r, err := BuildPlainResourceFromCitation(citation)
	require.NoError(t, err)
	require.IsType(t, PlainResource{}, r)
	p := r.(PlainResource)
	assert.Equal(t, u, p.FileURL)
	lr, _ := NewLineRange(3, 7)
	assert.Equal(t, lr, p.LineRange)
}

func TestBuildPlainResourceFromCitationError(t *testing.T) {
	citation := Citation{
		URL:   nil,
		Extra: "bad",
	}
	_, err := BuildPlainResourceFromCitation(citation)
	assert.Error(t, err)
}

func TestPlainResource(t *testing.T) {
	u, err := url.Parse("http://plain.com/path")
	require.NoError(t, err)
	lr, err := NewLineRange(3, 7)
	require.NoError(t, err)
	p := PlainResource{
		FileURL:   u,
		LineRange: lr,
	}

	assert.Equal(t, u, p.URL())
	assert.Equal(t, Citation{URL: u, Extra: "3-7"}, p.Cite())
	assert.Equal(t, lr, p.Lines())
}
