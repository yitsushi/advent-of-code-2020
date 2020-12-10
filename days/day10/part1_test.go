package day10_test

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2020/days/day10"
	"github.com/yitsushi/go-misskey/test"
)

func TestSolver_Part1(t *testing.T) {
	example, err := test.LoadFixture("example-short")
	if !assert.NoError(t, err) {
		return
	}

	day := day10.Solver{}

	day.SetInput(ioutil.NopCloser(bytes.NewReader(example)))

	out, err := day.Part1()

	assert.NoError(t, err)
	assert.Equal(t, "35", out)
}

func TestSolver_Part1_long(t *testing.T) {
	example, err := test.LoadFixture("example-long")
	if !assert.NoError(t, err) {
		return
	}

	day := day10.Solver{}

	day.SetInput(ioutil.NopCloser(bytes.NewReader(example)))

	out, err := day.Part1()

	assert.NoError(t, err)
	assert.Equal(t, "220", out)
}
