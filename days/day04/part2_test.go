package day04_test

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2020/days/day04"
)

func TestSolver_Part2(t *testing.T) {
	day := day04.Solver{}

	day.SetInput(ioutil.NopCloser(strings.NewReader(exampleInput)))

	out, err := day.Part2()

	assert.NoError(t, err)
	assert.Equal(t, "2", out)
}

func TestSolver_Part2_allInvalid(t *testing.T) {
	day := day04.Solver{}

	day.SetInput(ioutil.NopCloser(strings.NewReader(allInvalid)))

	out, err := day.Part2()

	assert.NoError(t, err)
	assert.Equal(t, "0", out)
}

func TestSolver_Part2_allValid(t *testing.T) {
	day := day04.Solver{}

	day.SetInput(ioutil.NopCloser(strings.NewReader(allValid)))

	out, err := day.Part2()

	assert.NoError(t, err)
	assert.Equal(t, "4", out)
}
