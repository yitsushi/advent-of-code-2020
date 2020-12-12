package day12_test

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2020/days/day12"
	"github.com/yitsushi/advent-of-code-2020/test"
)

func TestSolver_Part2_example(t *testing.T) {
	example, err := test.LoadFixture("example")
	if !assert.NoError(t, err) {
		return
	}

	day := day12.Solver{}

	day.SetInput(ioutil.NopCloser(bytes.NewReader(example)))

	out, err := day.Part2()

	assert.NoError(t, err)
	assert.Equal(t, "286", out)
}

func TestSolver_Part2_example2(t *testing.T) {
	example, err := test.LoadFixture("example2")
	if !assert.NoError(t, err) {
		return
	}

	day := day12.Solver{}

	day.SetInput(ioutil.NopCloser(bytes.NewReader(example)))

	out, err := day.Part2()

	assert.NoError(t, err)
	assert.Equal(t, "691", out)
}
