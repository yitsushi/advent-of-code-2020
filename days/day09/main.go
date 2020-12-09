package day09

import (
	"fmt"

	"github.com/yitsushi/advent-of-code-2020/pkg/perf"
	"github.com/yitsushi/advent-of-code-2020/pkg/puzzle"
)

const (
	realWindow = 25
)

// Part1 for this day.
func (d *Solver) Part1() (string, error) {
	defer perf.Duration(perf.Track("Part1"))

	invalid, err := d.findInvalid(realWindow)

	switch err.(type) {
	case PairNotFound:
		return fmt.Sprintf("%d", invalid), nil
	default:
		return "0", puzzle.NoSolution{}
	}
}

// Part2 for this day.
func (d *Solver) Part2() (string, error) {
	defer perf.Duration(perf.Track("Part2"))

	invalid, _ := d.findInvalid(realWindow)
	weakness, err := d.findWeakness(invalid)

	return fmt.Sprintf("%d", weakness), err
}
