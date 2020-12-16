package day16_test

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2020/days/day16"
	"github.com/yitsushi/advent-of-code-2020/test"
)

func TestSolver_SetInput(t *testing.T) {
	example, err := test.LoadFixture("example")
	if !assert.NoError(t, err) {
		return
	}

	day := day16.Solver{}

	err = day.SetInput(ioutil.NopCloser(bytes.NewReader(example)))

	assert.NoError(t, err)
}

func TestSolver_SetInput_invalidRange(t *testing.T) {
	example, err := test.LoadFixture("invalid-range")
	if !assert.NoError(t, err) {
		return
	}

	day := day16.Solver{}

	err = day.SetInput(ioutil.NopCloser(bytes.NewReader(example)))

	assert.Error(t, err)
}

func TestSolver_SetInput_invalidList(t *testing.T) {
	example, err := test.LoadFixture("invalid-ticket-list")
	if !assert.NoError(t, err) {
		return
	}

	day := day16.Solver{}

	err = day.SetInput(ioutil.NopCloser(bytes.NewReader(example)))

	assert.Error(t, err)
}
