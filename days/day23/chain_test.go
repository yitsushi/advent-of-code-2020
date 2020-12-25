package day23_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2020/days/day23"
)

func TestChain(t *testing.T) {
	list := []int{9, 1, 2, 8, 7, 3, 4, 6, 5}

	chain := day23.NewChainOf(list, 15)

	assert.Equal(t, 7, chain.Nth(4).Value)
	assert.Equal(t, 13, chain.Nth(12).Value)
	assert.Equal(t, 9, chain.Nth(30).Value)
	assert.Nil(t, chain.Find(50))
	assert.NotNil(t, chain.Find(6))
	assert.Equal(t, 6, chain.Find(6).Value)
}
