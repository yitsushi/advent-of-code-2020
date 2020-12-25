package day25_test

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2020/days/day25"
	"github.com/yitsushi/advent-of-code-2020/test"
)

func TestSolver_Part2(t *testing.T) {
	example, err := test.LoadFixture("example")
	if !assert.NoError(t, err) {
		return
	}

	day := day25.Solver{}

	err = day.SetInput(ioutil.NopCloser(bytes.NewReader(example)))
	if !assert.NoError(t, err) {
		return
	}

	out, err := day.Part2()

	assert.NoError(t, err)
	assert.Equal(t, "Happy Yule!", out)
}
