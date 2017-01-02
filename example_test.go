package cite

import (
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExample(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Setup processor
	builders := []ResourceBuilder{
		BuildGithubResourceFromCitation,
	}
	processor := NewProcessor(builders)
	processor.AddHandler("insert", InsertHandler)

	// Read input
	f, err := os.Open("example/example.go.pre")
	require.NoError(t, err)
	defer f.Close()

	src := ParseCode(f)

	// Mock HTTP call
	data, err := ioutil.ReadFile("example/grinch.txt")
	require.NoError(t, err)
	httpmock.RegisterResponder(
		http.MethodGet,
		"https://github.com/mmcloughlin/cite/raw/master/example/grinch.txt",
		httpmock.NewBytesResponder(200, data),
	)

	// Process it
	src, err = processor.Process(src)
	require.NoError(t, err)

	// Compare to expected
	expect, err := ioutil.ReadFile("example/example.go")
	require.NoError(t, err)

	assert.Equal(t, expect, []byte(src.String()))
}
