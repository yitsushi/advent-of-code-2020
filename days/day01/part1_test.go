package day01_test

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2020/days/day01"
)

func TestSolver_Part1(t *testing.T) {
	day := day01.Solver{}

	day.SetInput(ioutil.NopCloser(strings.NewReader("123\n1572\n885\n448")))

	out, err := day.Part1()

	assert.NoError(t, err)
	assert.Equal(t, "704256", out)
}

func TestSolver_Part1_noSolution(t *testing.T) {
	day := day01.Solver{}

	day.SetInput(ioutil.NopCloser(strings.NewReader("123\n543\n885")))

	out, err := day.Part1()

	assert.Error(t, err)
	assert.Equal(t, "", out)
}
