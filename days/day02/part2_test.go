package day02_test

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2020/days/day02"
)

func TestSolver_Part2(t *testing.T) {
	day := day02.Solver{}

	day.SetInput(ioutil.NopCloser(strings.NewReader(exampleInput)))

	out, err := day.Part2()

	assert.NoError(t, err)
	assert.Equal(t, "1", out)
}
