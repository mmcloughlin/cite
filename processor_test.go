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
	assert.Equal(t, "https://www.rfc-editor.org/rfc/rfc918.txt", dir.Citation.URL.String())
	assert.Equal(t, "13-16", dir.Citation.Extra)
}

func TestParseDirectiveNoExtra(t *testing.T) {
	line := "Action: http://google.org"
	dir, err := ParseDirective(line)
	require.NoError(t, err)
	assert.Equal(t, "Action", dir.ActionRaw)
	assert.Equal(t, "http://google.org", dir.Citation.URL.String())
	assert.Equal(t, "", dir.Citation.Extra)
}

func TestParseDirectiveGarbage(t *testing.T) {
	line := "jk lol"
	dir, err := ParseDirective(line)
	assert.Nil(t, dir)
	assert.NoError(t, err)
}

func TestDirectiveAction(t *testing.T) {
	dir := Directive{
		ActionRaw: "MiXeDCAsE",
	}
	assert.Equal(t, "mixedcase", dir.Action())
}
