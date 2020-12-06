package day06_test

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2020/days/day06"
)

func TestSolver_SetInput(t *testing.T) {
	day := day06.Solver{}
	err := day.SetInput(ioutil.NopCloser(strings.NewReader(exampleInput)))

	assert.NoError(t, err)
}
