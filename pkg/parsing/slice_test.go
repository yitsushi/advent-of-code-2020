package parsing_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2020/pkg/parsing"
)

func TestInt64Slice(t *testing.T) {
	input := "12,32,55,23"
	expected := []int64{12, 32, 55, 23}
	got, err := parsing.Int64Slice(input, ",")

	assert.NoError(t, err)
	assert.Equal(t, expected, got)
}

func TestInt64Slice_error(t *testing.T) {
	input := "12,asd,55,23"
	expected := []int64{12}
	got, err := parsing.Int64Slice(input, ",")

	assert.Error(t, err)
	assert.Equal(t, expected, got)
}
