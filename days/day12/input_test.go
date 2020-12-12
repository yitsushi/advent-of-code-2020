package day12_test

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2020/days/day12"
	"github.com/yitsushi/advent-of-code-2020/test"
)

func TestSolver_SetInput(t *testing.T) {
	example, err := test.LoadFixture("example")
	if !assert.NoError(t, err) {
		return
	}

	day := day12.Solver{}

	day.SetInput(ioutil.NopCloser(bytes.NewReader(example)))

	assert.NoError(t, err)
}

func TestSolver_SetInput_invalidCommand(t *testing.T) {
	day := day12.Solver{}
	err := day.SetInput(ioutil.NopCloser(strings.NewReader("K44")))

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "K44")
}

func TestSolver_SetInput_noValue(t *testing.T) {
	day := day12.Solver{}
	err := day.SetInput(ioutil.NopCloser(strings.NewReader("N")))

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "N")
}
