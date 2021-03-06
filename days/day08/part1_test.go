package day08_test

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2020/days/day08"
)

func TestSolver_Part1(t *testing.T) {
	day := day08.Solver{}

	day.SetInput(ioutil.NopCloser(strings.NewReader(exampleInput)))

	out, err := day.Part1()

	assert.NoError(t, err)
	assert.Equal(t, "5", out)
}

func TestSolver_Part1_noSolution(t *testing.T) {
	day := day08.Solver{}

	day.SetInput(ioutil.NopCloser(strings.NewReader("nop 0")))

	out, err := day.Part1()

	assert.Error(t, err)
	assert.Equal(t, "", out)
}
