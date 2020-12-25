package day23_test

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2020/days/day23"
)

func TestSolver_Part2(t *testing.T) {
	day := day23.Solver{}

	err := day.SetInput(ioutil.NopCloser(strings.NewReader("389125467")))
	if !assert.NoError(t, err) {
		return
	}

	out, err := day.Part2()

	assert.NoError(t, err)
	assert.Equal(t, "149245887792", out)
}
