package cite

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGithubResourceCite(t *testing.T) {
	lines, _ := NewLineRange(3, 10)
	res := GithubResource{
		User:       "mmcloughlin",
		Repository: "geohash",
		GitRef:     "master",
		Path:       "LICENSE",
		LineRange:  lines,
	}

	citation := res.Cite()
	expect := "https://github.com/mmcloughlin/geohash/blob/master/LICENSE#L3-L10"
	assert.Equal(t, expect, citation.URL.String())
	assert.Equal(t, "", citation.Extra)
}

//func TestReferenceStringRawURL(t *testing.T) {
//	ref := Reference{
//		User:       "a",
//		Repository: "b",
//		GitRef:     "master",
//		Path:       "path/to/file",
//		Lines:      NewLineRange(42, 103),
//	}
//	expect := "https://github.com/a/b/raw/master/path/to/file"
//	assert.Equal(t, expect, ref.RawFileURL().String())
//}

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
