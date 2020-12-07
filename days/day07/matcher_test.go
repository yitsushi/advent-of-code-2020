package day07_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2020/days/day07"
)

type matcherTestCase struct {
	Sentence string
	Expected day07.Rule
}

func TestMatcher_example1(t *testing.T) {
	testCases := []matcherTestCase{
		{
			Sentence: "light red bags contain 1 bright white bag, 2 muted yellow bags.",
			Expected: day07.Rule{
				Bag: day07.Bag{
					Tone:  "light",
					Color: "red",
					Name:  "light red",
				},
				Contains: []day07.ContentStatement{
					{Bag: day07.Bag{Tone: "bright", Color: "white", Name: "bright white"}, Count: 1},
					{Bag: day07.Bag{Tone: "muted", Color: "yellow", Name: "muted yellow"}, Count: 2},
				},
			},
		},
		{
			Sentence: "vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
			Expected: day07.Rule{
				Bag: day07.Bag{
					Tone:  "vibrant",
					Color: "plum",
					Name:  "vibrant plum",
				},
				Contains: []day07.ContentStatement{
					{Bag: day07.Bag{Tone: "faded", Color: "blue", Name: "faded blue"}, Count: 5},
					{Bag: day07.Bag{Tone: "dotted", Color: "black", Name: "dotted black"}, Count: 6},
				},
			},
		},
		{
			Sentence: "bright white bags contain 1 shiny gold bag.",
			Expected: day07.Rule{
				Bag: day07.Bag{
					Tone:  "bright",
					Color: "white",
					Name:  "bright white",
				},
				Contains: []day07.ContentStatement{
					{Bag: day07.Bag{Tone: "shiny", Color: "gold", Name: "shiny gold"}, Count: 1},
				},
			},
		},
	}

	for _, testCase := range testCases {
		rule := day07.Matcher(testCase.Sentence)

		assert.EqualValues(t, testCase.Expected, rule)
	}
}
