package day13

import (
	"fmt"

	amath "github.com/yitsushi/advent-of-code-2020/pkg/math"
	"github.com/yitsushi/advent-of-code-2020/pkg/perf"
)

// Part1 for this day.
func (d *Solver) Part1() (string, error) {
	defer perf.Duration(perf.Track("Part1"))

	var (
		minID    int64 = -1
		waitTime int64 = d.earliest
	)

	for _, bus := range d.busList {
		if bus == -1 {
			continue
		}

		_, r := amath.DivMod(d.earliest, bus)
		wait := bus - r

		if wait < waitTime {
			minID = bus
			waitTime = wait
		}
	}

	return fmt.Sprintf("%d", minID*waitTime), nil
}

// Part2 for this day.
func (d *Solver) Part2() (string, error) {
	defer perf.Duration(perf.Track("Part2"))

	var (
		ts   int64 = 0
		step int64 = 1
	)

	// It could be some lcm and gcd magic,
	// but all of them are prime.
	// At least in my input.
	for idx, bus := range d.busList {
		if bus == -1 {
			continue
		}

		for (ts+int64(idx))%bus != 0 {
			ts += step
		}

		step *= bus
	}

	return fmt.Sprintf("%d", ts), nil
}
