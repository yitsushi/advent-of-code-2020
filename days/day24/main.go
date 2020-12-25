package day24

import (
	"strconv"

	"github.com/yitsushi/advent-of-code-2020/pkg/perf"
)

// Part1 for this day.
func (d *Solver) Part1() (string, error) {
	defer perf.Duration(perf.Track("Part1"))

	grid := NewGrid()

	for _, directions := range d.instructions {
		grid.FlipFromInstruction(directions)
	}

	return strconv.Itoa(grid.Count(Black)), nil
}

// Part2 for this day.
func (d *Solver) Part2() (string, error) {
	defer perf.Duration(perf.Track("Part2"))

	grid := NewGrid()

	for _, directions := range d.instructions {
		grid.FlipFromInstruction(directions)
	}

	for day := 0; day < 100; day++ {
		grid.OneDayLater()
	}

	return strconv.Itoa(grid.Count(Black)), nil
}
