package day03_test

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2020/days/day03"
)

func TestSolver_Part1(t *testing.T) {
	day := day03.Solver{}

	day.SetInput(ioutil.NopCloser(strings.NewReader(exampleInput)))

	out, err := day.Part1()

	assert.NoError(t, err)
	assert.Equal(t, "7", out)
}
