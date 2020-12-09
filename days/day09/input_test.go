package day09_test

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2020/days/day09"
	"github.com/yitsushi/advent-of-code-2020/test"
)

func TestSolver_SetInput(t *testing.T) {
	example, err := test.LoadFixture("real")
	if !assert.NoError(t, err) {
		return
	}

	day := day09.Solver{}
	err = day.SetInput(ioutil.NopCloser(bytes.NewReader(example)))

	assert.NoError(t, err)
}
