package day15_test

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2020/days/day15"
)

func TestSolver_Part2(t *testing.T) {
	testCases := map[string]string{
		"0,3,6": "175594",
		"1,3,2": "2578",
		"2,1,3": "3544142",
		"1,2,3": "261214",
		"2,3,1": "6895259",
		"3,2,1": "18",
		"3,1,2": "362",
	}

	for input, expected := range testCases {
		day := day15.Solver{}

		err := day.SetInput(ioutil.NopCloser(strings.NewReader(input)))
		if !assert.NoError(t, err) {
			continue
		}

		out, err := day.Part2()

		assert.NoError(t, err)
		assert.Equal(t, expected, out)

		break
	}
}

func TestSolver_Part2_real(t *testing.T) {
	day := day15.Solver{}

	err := day.SetInput(ioutil.NopCloser(strings.NewReader("14,1,17,0,3,20")))
	if !assert.NoError(t, err) {
		return
	}

	out, err := day.Part2()

	assert.NoError(t, err)
	assert.Equal(t, "6428", out)
}
