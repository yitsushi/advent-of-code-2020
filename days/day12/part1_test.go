package day12_test

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2020/days/day12"
)

func TestSolver_Part1(t *testing.T) {
	day := day12.Solver{}

	day.SetInput(ioutil.NopCloser(strings.NewReader("")))

	out, err := day.Part1()

	assert.Error(t, err)
	assert.Equal(t, "", out)
}

func TestSolver_Part1_noSolution(t *testing.T) {
	day := day12.Solver{}

	day.SetInput(ioutil.NopCloser(strings.NewReader("")))

	out, err := day.Part1()

	assert.Error(t, err)
	assert.Equal(t, "", out)
}
