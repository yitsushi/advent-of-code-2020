package day11_test

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2020/days/day11"
)

func TestSolver_SetInput(t *testing.T) {
	day := day11.Solver{}
	err := day.SetInput(ioutil.NopCloser(strings.NewReader("##..#\nL##LL.#")))

	assert.NoError(t, err)
}

func TestSolver_SetInput_invalid(t *testing.T) {
	day := day11.Solver{}
	err := day.SetInput(ioutil.NopCloser(strings.NewReader("##$.#\nL##LL.#")))

	assert.Error(t, err)
	assert.Equal(t, "Unknown character: $", err.Error())
}
