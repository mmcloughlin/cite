package cite

import (
	"bytes"
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

// TestParseDirectiveBadURL tests the case where the URL regular expression
// matches but the extracted URL does not parse. This should be a rare case,
// but it is possible to find examples ("http%") being the most compact.
func TestParseDirectiveBadURL(t *testing.T) {
	line := "Action: http%"
	dir, err := ParseDirective(line)
	assert.Error(t, err)
	assert.Nil(t, dir)
}

func TestDirectiveAction(t *testing.T) {
	dir := Directive{
		ActionRaw: "MiXeDCAsE",
	}
	assert.Equal(t, "mixedcase", dir.Action())
}

func BuildSingleCommentSource(line string) Source {
	comment := "// " + line
	code := bytes.NewReader([]byte(comment))
	return ParseCode(code)
}

func TestProcessLinesParseDirectiveError(t *testing.T) {
	p := NewProcessor(nil)
	src := BuildSingleCommentSource("Action: http%")
	_, err := p.Process(src)
	assert.Error(t, err)
}

func TestProcessLinesUnknownResource(t *testing.T) {
	p := NewProcessor(nil)
	src := BuildSingleCommentSource("Action: http://unknown.com")
	_, err := p.Process(src)
	assert.Equal(t, ErrUnknownResource, err)
}

func TestProcessLinesBadResource(t *testing.T) {
	builders := []ResourceBuilder{BuildGithubResourceFromCitation}
	p := NewProcessor(builders)
	src := BuildSingleCommentSource("Action: http://github.com/bad/path")
	_, err := p.Process(src)
	assert.Error(t, err)
}

func TestProcessLinesErrUnknownAction(t *testing.T) {
	builders := []ResourceBuilder{BuildPlainResourceFromCitation}
	p := NewProcessor(builders)
	src := BuildSingleCommentSource("Action: http://website.com/doc.txt (1-2)")
	_, err := p.Process(src)
	assert.Equal(t, ErrUnknownAction, err)
}

func TestProcessLinesHandlerError(t *testing.T) {
	builders := []ResourceBuilder{BuildPlainResourceFromCitation}
	p := NewProcessor(builders)
	p.AddHandler("action", func(_ Resource, _ []string) ([]string, []string, error) {
		return nil, nil, assert.AnError
	})
	src := BuildSingleCommentSource("Action: http://website.com/doc.txt (1-2)")
	_, err := p.Process(src)
	assert.Equal(t, assert.AnError, err)
}

// TODO error getting resource (eg bad github ref)
// TODO didnt find any resource
// TODO no handler for action
// TODO error in handler
