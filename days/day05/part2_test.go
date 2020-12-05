package day05_test

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2020/days/day05"
	"github.com/yitsushi/advent-of-code-2020/test"
)

func TestSolver_Part2(t *testing.T) {
	example, err := test.LoadFixture("real")
	if !assert.NoError(t, err) {
		return
	}

	day := day05.Solver{}

	err = day.SetInput(ioutil.NopCloser(bytes.NewReader(example)))
	if !assert.NoError(t, err) {
		return
	}

	out, err := day.Part2()

	assert.NoError(t, err)
	assert.Equal(t, "657", out)
}

func TestSolver_Part2_noSolution(t *testing.T) {
	example, err := test.LoadFixture("no-solution")
	if !assert.NoError(t, err) {
		return
	}

	day := day05.Solver{}

	err = day.SetInput(ioutil.NopCloser(bytes.NewReader(example)))
	if !assert.NoError(t, err) {
		return
	}

	out, err := day.Part2()

	assert.Error(t, err)
	assert.Equal(t, "", out)
}
