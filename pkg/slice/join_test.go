package slice_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2020/pkg/slice"
)

func TestJoinInt(t *testing.T) {
	input := []int{5, 2, 6, 8}
	expected := "5;2;6;8"

	assert.Equal(t, expected, slice.JoinInt(input, ";"))
}
