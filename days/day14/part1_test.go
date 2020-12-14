package day14_test

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2020/days/day14"
	"github.com/yitsushi/advent-of-code-2020/test"
)

func TestSolver_Part1(t *testing.T) {
	example, err := test.LoadFixture("example")
	if !assert.NoError(t, err) {
		return
	}

	day := day14.Solver{}

	day.SetInput(ioutil.NopCloser(bytes.NewReader(example)))

	out, err := day.Part1()

	assert.NoError(t, err)
	assert.Equal(t, "165", out)
}

func TestSolver_Part1_real(t *testing.T) {
	example, err := test.LoadFixture("input")
	if !assert.NoError(t, err) {
		return
	}

	day := day14.Solver{}

	day.SetInput(ioutil.NopCloser(bytes.NewReader(example)))

	out, err := day.Part1()

	assert.NoError(t, err)
	assert.Equal(t, "11501064782628", out)
}
