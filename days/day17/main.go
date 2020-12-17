package day17

import (
	"fmt"

	"github.com/yitsushi/advent-of-code-2020/pkg/perf"
	"github.com/yitsushi/advent-of-code-2020/pkg/puzzle"
)

// Part1 for this day.
func (d *Solver) Part1() (string, error) {
	defer perf.Duration(perf.Track("Part1"))

	d.simulate()

	return fmt.Sprintf("%d", len(d.space.SelectNodes(true))), nil
}

// Part2 for this day.
func (d *Solver) Part2() (string, error) {
	defer perf.Duration(perf.Track("Part2"))

	return "", puzzle.NotImplemented{}
}
