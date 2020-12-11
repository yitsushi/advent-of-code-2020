package day11

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/yitsushi/advent-of-code-2020/pkg/perf"
)

const (
	occupationThresholdPart1 = 3
	occupationThresholdPart2 = 4
)

// Part1 for this day.
func (d *Solver) Part1() (string, error) {
	defer perf.Duration(perf.Track("Part1"))

	for {
		changes := d.ferry.Cycle(occupationThresholdPart1, false)
		if changes == 0 {
			break
		}
	}

	logrus.Infof("Number of occupied seats: %d", d.ferry.Count(Occupied))

	return fmt.Sprintf("%d", d.ferry.Count(Occupied)), nil
}

// Part2 for this day.
func (d *Solver) Part2() (string, error) {
	defer perf.Duration(perf.Track("Part2"))

	for {
		changes := d.ferry.Cycle(occupationThresholdPart2, true)
		if changes == 0 {
			break
		}
	}

	logrus.Infof("Number of occupied seats: %d", d.ferry.Count(Occupied))

	return fmt.Sprintf("%d", d.ferry.Count(Occupied)), nil
}
