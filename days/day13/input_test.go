package day13_test

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2020/days/day13"
)

func TestSolver_SetInput(t *testing.T) {
	day := day13.Solver{}
	err := day.SetInput(ioutil.NopCloser(strings.NewReader("321\n23,43,12,x,44")))

	assert.NoError(t, err)
}
