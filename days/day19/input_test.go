package day19_test

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2020/days/day19"
	"github.com/yitsushi/advent-of-code-2020/test"
)

func TestSolver_SetInput(t *testing.T) {
	t.Skip()

	example, err := test.LoadFixture("example")
	if !assert.NoError(t, err) {
		return
	}

	day := day19.Solver{}

	day.SetInput(ioutil.NopCloser(bytes.NewReader(example)))

	assert.NoError(t, err)
}
