package day23_test

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2020/days/day23"
)

func TestSolver_Part1(t *testing.T) {
	day := day23.Solver{}

	err := day.SetInput(ioutil.NopCloser(strings.NewReader("389125467")))
	if !assert.NoError(t, err) {
		return
	}

	out, err := day.Part1()

	assert.NoError(t, err)
	assert.Equal(t, "67384529", out)
}

func TestSolver_Part1_real(t *testing.T) {
	day := day23.Solver{}

	err := day.SetInput(ioutil.NopCloser(strings.NewReader("219347865")))
	if !assert.NoError(t, err) {
		return
	}

	out, err := day.Part1()

	assert.NoError(t, err)
	assert.Equal(t, "36472598", out)
}
