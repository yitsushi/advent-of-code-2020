package day07_test

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2020/days/day07"
)

func TestSolver_Part2_example1(t *testing.T) {
	day := day07.Solver{}

	day.SetInput(ioutil.NopCloser(strings.NewReader(exampleInput)))

	out, err := day.Part2()

	assert.NoError(t, err)
	assert.Equal(t, "32", out)
}

func TestSolver_Part2_example2(t *testing.T) {
	day := day07.Solver{}

	day.SetInput(ioutil.NopCloser(strings.NewReader(anotherExample)))

	out, err := day.Part2()

	assert.NoError(t, err)
	assert.Equal(t, "126", out)
}
