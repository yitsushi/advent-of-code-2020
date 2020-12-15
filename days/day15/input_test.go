package day15_test

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2020/days/day15"
)

func TestSolver_SetInput(t *testing.T) {
	day := day15.Solver{}
	err := day.SetInput(ioutil.NopCloser(strings.NewReader("1,2,3")))

	assert.NoError(t, err)
}

func TestSolver_SetInput_empty(t *testing.T) {
	day := day15.Solver{}
	err := day.SetInput(ioutil.NopCloser(strings.NewReader("")))

	assert.NoError(t, err)
}

func TestSolver_SetInput_invalid(t *testing.T) {
	day := day15.Solver{}
	err := day.SetInput(ioutil.NopCloser(strings.NewReader("k,j,s")))

	assert.Error(t, err)
}
