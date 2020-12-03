package day03_test

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2020/days/day03"
)

func TestSolver_Part2(t *testing.T) {
	day := day03.Solver{}

	day.SetInput(ioutil.NopCloser(strings.NewReader(exampleInput)))

	out, err := day.Part2()

	assert.NoError(t, err)
	assert.Equal(t, "336", out)
}

func TestSolver_Part2_invalidSquare(t *testing.T) {
	day := day03.Solver{}

	day.SetInput(ioutil.NopCloser(strings.NewReader(".....#\n&&&&&&")))

	out, err := day.Part2()

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "&")
	assert.Equal(t, "", out)
}
