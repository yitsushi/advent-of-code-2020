package day24_test

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2020/days/day24"
	"github.com/yitsushi/advent-of-code-2020/test"
)

func TestSolver_SetInput(t *testing.T) {
	example, err := test.LoadFixture("example")
	if !assert.NoError(t, err) {
		return
	}

	day := day24.Solver{}
	err = day.SetInput(ioutil.NopCloser(bytes.NewReader(example)))

	assert.NoError(t, err)
}

func TestSolver_SetInput_invalid(t *testing.T) {
	example, err := test.LoadFixture("invalid")
	if !assert.NoError(t, err) {
		return
	}

	day := day24.Solver{}
	err = day.SetInput(ioutil.NopCloser(bytes.NewReader(example)))

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unknown direction")
}
