package day01

import (
	"fmt"

	"github.com/yitsushi/advent-of-code-2020/pkg/perf"
	"github.com/yitsushi/advent-of-code-2020/pkg/puzzle"
)

const (
	target int64 = 2020
)

// Solver is the main solver.
type Solver struct {
	input []int64
}

// Part1 for this day.
func (d *Solver) Part1() (string, error) {
	defer perf.Duration(perf.Track("Part1"))

	for idx, v1 := range d.input {
		for _, v2 := range d.input[idx+1:] {
			if v1+v2 == target {
				return fmt.Sprintf("%d", v1*v2), nil
			}
		}
	}

	return "", puzzle.NoSolution{}
}

// Part2 for this day.
func (d *Solver) Part2() (string, error) {
	defer perf.Duration(perf.Track("Part2"))

	for idx, v1 := range d.input {
		for jdx, v2 := range d.input[idx+1:] {
			for _, v3 := range d.input[jdx+1:] {
				if v1+v2+v3 == target {
					return fmt.Sprintf("%d", v1*v2*v3), nil
				}
			}
		}
	}

	return "", puzzle.NotImplemented{}
}
