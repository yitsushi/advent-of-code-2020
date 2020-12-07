package day07_test

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2020/days/day07"
)

func TestSolver_Part1(t *testing.T) {
	day := day07.Solver{}

	day.SetInput(ioutil.NopCloser(strings.NewReader(exampleInput)))

	out, err := day.Part1()

	assert.NoError(t, err)
	assert.Equal(t, "4", out)
}
