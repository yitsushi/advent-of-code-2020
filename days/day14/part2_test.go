package day14_test

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2020/days/day14"
	"github.com/yitsushi/advent-of-code-2020/test"
)

func TestSolver_Part2(t *testing.T) {
	example, err := test.LoadFixture("example2")
	if !assert.NoError(t, err) {
		return
	}

	day := day14.Solver{}

	day.SetInput(ioutil.NopCloser(bytes.NewReader(example)))

	out, err := day.Part2()

	assert.NoError(t, err)
	assert.Equal(t, "208", out)
}

func TestSolver_Part2_real(t *testing.T) {
	example, err := test.LoadFixture("input")
	if !assert.NoError(t, err) {
		return
	}

	day := day14.Solver{}

	day.SetInput(ioutil.NopCloser(bytes.NewReader(example)))

	out, err := day.Part2()

	assert.NoError(t, err)
	assert.Equal(t, "5142195937660", out)
}
