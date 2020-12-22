package parsing_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2020/pkg/parsing"
)

func TestInt64Range(t *testing.T) {
	input := "12-34"
	expectedFrom := int64(12)
	expectedTo := int64(34)
	gotFrom, gotTo, err := parsing.Int64Range(input, "-")

	assert.NoError(t, err)
	assert.Equal(t, expectedFrom, gotFrom)
	assert.Equal(t, expectedTo, gotTo)
}

func TestInt64Range_errorFrom(t *testing.T) {
	input := "as-34"
	expectedFrom := int64(0)
	expectedTo := int64(0)
	gotFrom, gotTo, err := parsing.Int64Range(input, "-")

	assert.Error(t, err)
	assert.Equal(t, expectedFrom, gotFrom)
	assert.Equal(t, expectedTo, gotTo)
}

func TestInt64Range_errorTo(t *testing.T) {
	input := "12-as"
	expectedFrom := int64(12)
	expectedTo := int64(0)
	gotFrom, gotTo, err := parsing.Int64Range(input, "-")

	assert.Error(t, err)
	assert.Equal(t, expectedFrom, gotFrom)
	assert.Equal(t, expectedTo, gotTo)
}
