package cite

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseLineRangeSingle(t *testing.T) {
	lr, err := ParseLineRange("7")
	require.NoError(t, err)
	assert.Equal(t, NewSingleLine(7), lr)
}

func TestParseLineRange(t *testing.T) {
	lr, err := ParseLineRange("7-13")
	require.NoError(t, err)
	expect, err := NewLineRange(7, 13)
	require.NoError(t, err)
	assert.Equal(t, expect, lr)
}

func TestParseLineRangeErrors(t *testing.T) {
	cases := []string{
		"1-2-3",
		"idk",
		"3-jk",
		"7-3",
	}
	for _, s := range cases {
		_, err := ParseLineRange(s)
		assert.Error(t, err)
	}
}

func TestLineRangeString(t *testing.T) {
	lr, _ := NewLineRange(7, 13)
	assert.Equal(t, "7-13", lr.String())
}

func TestLineRangeStringSingle(t *testing.T) {
	lr := NewSingleLine(7)
	assert.Equal(t, "7", lr.String())
}
