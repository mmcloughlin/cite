package cite

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseDirective(t *testing.T) {
	line := "insert: https://www.rfc-editor.org/rfc/rfc918.txt (13-16)"
	dir, err := ParseDirective(line)
	require.NoError(t, err)
	assert.Equal(t, "insert", dir.ActionRaw)
	assert.Equal(t, "https://www.rfc-editor.org/rfc/rfc918.txt", dir.URL.String())
	assert.Equal(t, "13-16", dir.Extra)
}

func TestParseDirectiveNoExtra(t *testing.T) {
	line := "Action: http://google.org"
	dir, err := ParseDirective(line)
	require.NoError(t, err)
	assert.Equal(t, "Action", dir.ActionRaw)
	assert.Equal(t, "http://google.org", dir.URL.String())
	assert.Equal(t, "", dir.Extra)
}

func TestParseDirectiveGarbage(t *testing.T) {
	line := "jk lol"
	dir, err := ParseDirective(line)
	assert.Nil(t, dir)
	assert.NoError(t, err)
}
