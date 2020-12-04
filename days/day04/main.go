package day04

import (
	"fmt"

	"github.com/yitsushi/advent-of-code-2020/pkg/perf"
)

// Part1 for this day.
func (d *Solver) Part1() (string, error) {
	defer perf.Duration(perf.Track("Part1"))

	counter := 0

	for _, p := range d.input {
		if p.ValidateFake() {
			counter++
		}
	}

	return fmt.Sprintf("%d", counter), nil
}

// Part2 for this day.
func (d *Solver) Part2() (string, error) {
	defer perf.Duration(perf.Track("Part2"))

	counter := 0

	for _, p := range d.input {
		if p.Validate() {
			counter++
		}
	}

	return fmt.Sprintf("%d", counter), nil
}
