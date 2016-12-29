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

func TestLineRangeGetters(t *testing.T) {
	lr, err := NewLineRange(7, 13)
	require.NoError(t, err)
	assert.Equal(t, 7, lr.Start())
	assert.Equal(t, 13, lr.End())
	assert.Equal(t, 7, lr.NumLines())
}

func TestLineRangeString(t *testing.T) {
	lr, _ := NewLineRange(7, 13)
	assert.Equal(t, "7-13", lr.String())
}

func TestLineRangeStringSingle(t *testing.T) {
	lr := NewSingleLine(7)
	assert.Equal(t, "7", lr.String())
}

func TestLineRangeInclusion(t *testing.T) {
	lr, err := NewLineRange(7, 13)
	require.NoError(t, err)
	assert.False(t, lr.LineIncluded(6))
	assert.True(t, lr.LineIncluded(7))
	assert.True(t, lr.LineIncluded(10))
	assert.True(t, lr.LineIncluded(13))
	assert.False(t, lr.LineIncluded(14))
}
