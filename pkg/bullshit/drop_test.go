package bullshit_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2020/pkg/bullshit"
	"github.com/yitsushi/advent-of-code-2020/pkg/puzzle"
)

func TestDropErrorBoolean(t *testing.T) {
	assert.True(t, bullshit.DropErrorBoolean(true, puzzle.NoInputError{}))
	assert.False(t, bullshit.DropErrorBoolean(false, puzzle.NoInputError{}))
	assert.True(t, bullshit.DropErrorBoolean(true, nil))
	assert.False(t, bullshit.DropErrorBoolean(false, nil))
}

func TestDropErrorInt64(t *testing.T) {
	assert.Equal(t, int64(78), bullshit.DropErrorInt64(78, puzzle.NoInputError{}))
	assert.Equal(t, int64(99), bullshit.DropErrorInt64(99, nil))
}

func TestDropErrorUint64(t *testing.T) {
	assert.Equal(t, uint64(78), bullshit.DropErrorUint64(78, puzzle.NoInputError{}))
	assert.Equal(t, uint64(99), bullshit.DropErrorUint64(99, nil))
}
