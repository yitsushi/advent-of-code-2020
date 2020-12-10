package day10_test

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2020/days/day10"
)

func TestSolver_SetInput(t *testing.T) {
	day := day10.Solver{}
	err := day.SetInput(ioutil.NopCloser(strings.NewReader("12\n23\n1")))

	assert.NoError(t, err)
}

func TestSolver_SetInput_invalid(t *testing.T) {
	day := day10.Solver{}
	err := day.SetInput(ioutil.NopCloser(strings.NewReader("12\nasdsad\n44")))

	assert.Error(t, err)
}
