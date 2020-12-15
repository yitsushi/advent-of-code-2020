package day14_test

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2020/days/day14"
)

func TestSolver_SetInput(t *testing.T) {
	day := day14.Solver{}
	err := day.SetInput(ioutil.NopCloser(strings.NewReader("mask = xxx")))

	assert.NoError(t, err)
}

func TestSolver_SetInput_invalidIgnore(t *testing.T) {
	day := day14.Solver{}
	err := day.SetInput(ioutil.NopCloser(strings.NewReader("stuff = 123")))

	assert.NoError(t, err)
}
