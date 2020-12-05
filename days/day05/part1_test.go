package day05_test

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2020/days/day05"
	"github.com/yitsushi/advent-of-code-2020/test"
)

type testCase struct {
	Input  string
	Output string
}

func TestSolver_Part1(t *testing.T) {
	cases := []testCase{
		{Input: "FBFBBFFRLR", Output: "357"},
		{Input: "BFFFBBFRRR", Output: "567"},
		{Input: "FFFBBBFRRR", Output: "119"},
		{Input: "BBFFBBFRLL", Output: "820"},
	}

	for _, example := range cases {
		day := day05.Solver{}

		day.SetInput(ioutil.NopCloser(strings.NewReader(example.Input)))

		out, err := day.Part1()

		assert.NoError(t, err)
		assert.Equal(t, example.Output, out)
	}
}

func TestSolver_Part1_real(t *testing.T) {
	example, err := test.LoadFixture("real")
	if !assert.NoError(t, err) {
		return
	}

	day := day05.Solver{}

	err = day.SetInput(ioutil.NopCloser(bytes.NewReader(example)))
	if !assert.NoError(t, err) {
		return
	}

	out, err := day.Part1()

	assert.NoError(t, err)
	assert.Equal(t, "926", out)
}
