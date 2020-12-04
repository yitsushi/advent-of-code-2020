package day23_test

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2020/days/day23"
)

func TestSolver_Part2(t *testing.T) {
	day := day23.Solver{}

	day.SetInput(ioutil.NopCloser(strings.NewReader("")))

	out, err := day.Part2()

	assert.Error(t, err)
	assert.Equal(t, "", out)
}

func TestSolver_Part2_noSolution(t *testing.T) {
	day := day23.Solver{}

	day.SetInput(ioutil.NopCloser(strings.NewReader("")))

	out, err := day.Part2()

	assert.Error(t, err)
	assert.Equal(t, "", out)
}
