package day15_test

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2020/days/day15"
)

func TestSolver_Part1(t *testing.T) {
	testCases := map[string]string{
		"0,3,6": "436",
		"1,3,2": "1",
		"2,1,3": "10",
		"1,2,3": "27",
		"2,3,1": "78",
		"3,2,1": "438",
		"3,1,2": "1836",
	}

	for input, expected := range testCases {
		day := day15.Solver{}

		err := day.SetInput(ioutil.NopCloser(strings.NewReader(input)))
		if !assert.NoError(t, err) {
			continue
		}

		out, err := day.Part1()

		assert.NoError(t, err)
		assert.Equal(t, expected, out)

		break
	}
}
