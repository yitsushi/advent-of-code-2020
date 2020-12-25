package day23_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2020/days/day23"
)

func TestCup(t *testing.T) {
	list := []int{9, 1, 2, 8, 7, 3, 4, 6, 5}

	head := &day23.Cup{Value: list[0]}
	last := head

	for _, value := range list[1:] {
		new := &day23.Cup{Value: value}
		last.Next = new
		last = new
	}

	last.Next = head

	assert.Equal(t, 1, head.Next.Value)
	assert.Equal(t, 7, head.Nth(4).Value)
	assert.Equal(t, 8, head.Nth(12).Value)
	assert.Nil(t, head.Find(13))
	assert.NotNil(t, head.Find(6))
	assert.Equal(t, 6, head.Find(6).Value)
}
