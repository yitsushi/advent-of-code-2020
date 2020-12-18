package day18

import (
	"fmt"

	"github.com/yitsushi/advent-of-code-2020/pkg/perf"
)

// Part1 for this day.
func (d *Solver) Part1() (string, error) {
	defer perf.Duration(perf.Track("Part1"))

	var sum int64

	for _, str := range d.input {
		sum += calculateSimple(removeParentheses(str, calculateSimple))
	}

	return fmt.Sprintf("%d", sum), nil
}

// Part2 for this day.
func (d *Solver) Part2() (string, error) {
	defer perf.Duration(perf.Track("Part2"))

	var sum int64

	for _, str := range d.input {
		sum += calculateSimple(removeAddition(removeParentheses(str, calculateAdvanced)))
	}

	return fmt.Sprintf("%d", sum), nil
}
