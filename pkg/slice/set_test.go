package slice_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2020/pkg/slice"
)

func TestIntersectString(t *testing.T) {
	a := []string{"asd", "lkj", "qwe"}
	b := []string{"asd", "qwe", "poi"}
	c := []string{"asd", "qwe", "poi"}
	d := []string{"qwe", "gfd", "bvc"}
	expected := []string{"qwe"}

	assert.Equal(t, expected, slice.IntersectString(a, b, c, d))
	assert.Equal(t, a, slice.IntersectString(a))
	assert.Equal(t, []string{}, slice.IntersectString())
}
