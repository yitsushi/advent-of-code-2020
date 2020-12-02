package day02

import (
	"fmt"

	"github.com/yitsushi/advent-of-code-2020/pkg/perf"
)

// Solver is the main solver.
type Solver struct {
	input []entry
}

// Part1 for this day.
func (d *Solver) Part1() (string, error) {
	defer perf.Duration(perf.Track("Part1"))

	validPassword := 0

	for _, e := range d.input {
		found := 0

		for _, c := range e.Password {
			if e.Character == c {
				found += 1
			}
		}

		if e.Max >= found && found >= e.Min {
			validPassword += 1
		}
	}

	return fmt.Sprintf("%d", validPassword), nil
}

// Part2 for this day.
func (d *Solver) Part2() (string, error) {
	defer perf.Duration(perf.Track("Part2"))

	validPassword := 0

	for _, e := range d.input {
		first := e.Password[e.Min-1] == e.Character
		second := e.Password[e.Max-1] == e.Character

		if first != second {
			validPassword += 1
		}
	}

	return fmt.Sprintf("%d", validPassword), nil
}
