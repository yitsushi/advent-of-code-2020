package day08_test

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2020/days/day08"
)

func TestSolver_SetInput(t *testing.T) {
	day := day08.Solver{}
	err := day.SetInput(ioutil.NopCloser(strings.NewReader(exampleInput)))

	assert.NoError(t, err)
}

func TestSolver_SetInput_invalid(t *testing.T) {
	day := day08.Solver{}
	err := day.SetInput(ioutil.NopCloser(strings.NewReader("asdasdas +78")))

	assert.Error(t, err)
}
