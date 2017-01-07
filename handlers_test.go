package cite

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFormatSnippet(t *testing.T) {
	lines := []string{"Hello,", "World!"}
	expect := []string{"", "\tHello,", "\tWorld!", ""}
	assert.Equal(t, expect, FormatSnippet(lines))
}

func ExampleFormatSnippet() {
	snippet := FormatSnippet([]string{"Hello,", "World!"})
	for _, line := range snippet {
		fmt.Printf(">%s\n", line)
	}
	// Output:
	// >
	// >	Hello,
	// >	World!
	// >
}

func GenerateResource(t *testing.T) {
}

func TestInsertHandler(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	u, err := url.Parse("http://idk.com/doc.txt")
	require.NoError(t, err)

	lr, err := NewLineRange(2, 5)
	require.NoError(t, err)

	r := PlainResource{
		FileURL:   u,
		LineRange: lr,
	}

	data := "The\ncat\nin\nthe\nhat\nsat\non\nthe\nmat."

	httpmock.RegisterResponder(
		http.MethodGet,
		u.String(),
		httpmock.NewStringResponder(200, data),
	)

	lines := []string{
		fmt.Sprintf("Insert: %s", u.String()),
		"more",
		"lines",
	}
	insertion, rest, err := InsertHandler(r, lines)
	require.NoError(t, err)

	assert.Equal(t, []string{
		" Reference: http://idk.com/doc.txt (2-5)",
		"", "\tcat", "\tin", "\tthe", "\that", "",
	}, insertion)
	assert.Equal(t, []string{"more", "lines"}, rest)
}

func TestInsertHandlerError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	u, err := url.Parse("http://gone.com/doc.txt")
	require.NoError(t, err)

	r := &MockResource{}
	r.On("URL").Return(u).Once()

	httpmock.RegisterResponder(
		http.MethodGet,
		u.String(),
		httpmock.NewStringResponder(404, ""),
	)

	insertion, rest, err := InsertHandler(r, []string{})
	assert.Error(t, err)
	assert.Nil(t, insertion)
	assert.Nil(t, rest)
}

func TestSkipReferenceHandler(t *testing.T) {
	ls := &MockLineSelection{}
	ls.On("NumLines").Return(4).Once()

	r := &MockResource{}
	r.On("Lines").Return(ls).Once()

	ref := []string{
		" Reference: http://idk.com/doc.txt (2-5)",
		"", "\tcat", "\tin", "\tthe", "\that", "",
	}

	more := []string{"more", "lines"}

	lines := append(ref, more...)

	insertion, rest, err := SkipReferenceHandler(r, lines)

	require.NoError(t, err)
	assert.Equal(t, ref, insertion)
	assert.Equal(t, more, rest)
}

func TestSkipReferenceHandlerShort(t *testing.T) {
	ls := &MockLineSelection{}
	ls.On("NumLines").Return(2).Once()

	r := &MockResource{}
	r.On("Lines").Return(ls).Once()

	lines := []string{
		" Reference: http://idk.com/doc.txt (2-5)",
		"", "\tonly one line should be two", "",
	}

	_, _, err := SkipReferenceHandler(r, lines)
	assert.Error(t, err)
}
