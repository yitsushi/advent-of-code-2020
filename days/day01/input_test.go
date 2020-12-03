package day01_test

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2020/days/day01"
)

func TestSolver_SetInput(t *testing.T) {
	day := day01.Solver{}

	assert.NoError(t, day.SetInput(ioutil.NopCloser(strings.NewReader("123\n1572\n885\n448"))))
}

func TestSolver_SetInput_invalid(t *testing.T) {
	day := day01.Solver{}
	err := day.SetInput(ioutil.NopCloser(strings.NewReader("asdasd")))

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "\"asdasd\"")
}
