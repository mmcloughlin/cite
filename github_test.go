package cite

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
