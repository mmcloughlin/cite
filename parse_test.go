package cite

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func AssertParseCodeRoundTrip(t *testing.T, filename string) {
	data, err := ioutil.ReadFile(filename)
	require.NoError(t, err)

	src := ParseCode(bytes.NewBuffer(data))
	assert.Equal(t, string(data), src.String())
}

func TestParseCodeRoundTrip(t *testing.T) {
	gofiles, err := filepath.Glob("./*.go")
	require.NoError(t, err)
	for _, gofile := range gofiles {
		t.Log(gofile)
		AssertParseCodeRoundTrip(t, gofile)
	}
}

func ExampleCommentBlock_String() {
	c := CommentBlock{
		Leader: "/*leader*/ ",
		Lines: []string{
			"Hello,",
			"World!",
		},
	}
	fmt.Printf("%s", c)
	// Output:
	// /*leader*/ //Hello,
	// /*leader*/ //World!
}
