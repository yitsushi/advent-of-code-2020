package slice_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2020/pkg/slice"
)

func TestReverseString(t *testing.T) {
	type testCase struct {
		input    string
		expected string
	}

	testCases := []testCase{
		{input: "asdfghjkl", expected: "lkjhgfdsa"},
		{input: "123456789", expected: "987654321"},
	}

	for _, tc := range testCases {
		out := slice.ReverseString(tc.input)

		assert.Equal(t, tc.expected, out)
	}
}
