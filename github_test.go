package cite

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGithubResource(t *testing.T) {
	lines, _ := NewLineRange(3, 10)
	res := GithubResource{
		User:       "mmcloughlin",
		Repository: "geohash",
		GitRef:     "master",
		Path:       "LICENSE",
		LineRange:  lines,
	}

	expectRawURL := "https://github.com/mmcloughlin/geohash/raw/master/LICENSE"
	assert.Equal(t, expectRawURL, res.URL().String())

	citation := res.Cite()
	expect := "https://github.com/mmcloughlin/geohash/blob/master/LICENSE#L3-L10"
	assert.Equal(t, expect, citation.URL.String())
	assert.Equal(t, "", citation.Extra)

	assert.Equal(t, lines, res.Lines())
}

func TestLineRangeFragment(t *testing.T) {
	lr, _ := NewLineRange(7, 11)
	assert.Equal(t, "L7-L11", lineRangeFragment(lr))

	assert.Equal(t, "L7", lineRangeFragment(NewSingleLine(7)))

	assert.Panics(t, func() {
		lineRangeFragment(LineRange{
			start: 2,
			end:   1,
		})
	})
}

func TestParseLineRangeFragmentGarbage(t *testing.T) {
	_, err := parseLineRangeFragment("idk!?")
	assert.Error(t, err)
}

func TestParseLineRangeFragmentSingle(t *testing.T) {
	lines, err := parseLineRangeFragment("L42")
	assert.NoError(t, err)
	assert.Equal(t, NewSingleLine(42), lines)
}

func TestParseLineRangeFragmentRange(t *testing.T) {
	lines, err := parseLineRangeFragment("L3-L5")
	assert.NoError(t, err)
	expect, err := NewLineRange(3, 5)
	require.NoError(t, err)
	assert.Equal(t, expect, lines)
}

func BuildGithubResourceFromURL(urlStr string) (Resource, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		panic(err)
	}
	return BuildGithubResourceFromCitation(Citation{URL: u})
}

func TestBuildGithubResourceFromCitation(t *testing.T) {
	urlStr := "https://github.com/mmcloughlin/geohash/blob/master/LICENSE#L3-L10"
	res, err := BuildGithubResourceFromURL(urlStr)
	require.NoError(t, err)

	require.IsType(t, GithubResource{}, res)
	gh := res.(GithubResource)

	assert.Equal(t, "mmcloughlin", gh.User)
	assert.Equal(t, "geohash", gh.Repository)
	assert.Equal(t, "master", gh.GitRef)
	assert.Equal(t, "LICENSE", gh.Path)
	lr, _ := NewLineRange(3, 10)
	assert.Equal(t, lr, gh.LineRange)
}

func TestBuildGithubResourceFromCitationNotGithub(t *testing.T) {
	r, err := BuildGithubResourceFromURL("http://notgithub.com")
	assert.NoError(t, err)
	assert.Nil(t, r)
}

func TestBuildGithubResourceFromCitationErrors(t *testing.T) {
	cases := []string{
		"http://github.com",                                              // wrong scheme
		"https://github.com/short/path",                                  // short path
		"https://github.com/mmcloughlin/geohash/idk/master/LICENSE",      // missing "blob"
		"https://github.com/mmcloughlin/geohash/blob/master/LICENSE",     // no fragment
		"https://github.com/mmcloughlin/geohash/blob/master/LICENSE#bad", // bad fragment
	}
	for _, urlStr := range cases {
		r, err := BuildGithubResourceFromURL(urlStr)
		assert.Error(t, err)
		assert.Nil(t, r)
	}
}
