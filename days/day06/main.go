package day06

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/yitsushi/advent-of-code-2020/pkg/perf"
)

// Part1 for this day.
func (d *Solver) Part1() (string, error) {
	defer perf.Duration(perf.Track("Part1"))

	sumYes := 0

	for _, group := range d.input {
		logrus.Debug(group)
		sumYes += len(group.UniqueYesAnswers)
	}

	return fmt.Sprintf("%d", sumYes), nil
}

// Part2 for this day.
func (d *Solver) Part2() (string, error) {
	defer perf.Duration(perf.Track("Part2"))

	sumYes := 0

	for _, group := range d.input {
		// logrus.Debug(group)
		sumYes += len(group.UnanimousYes)
	}

	return fmt.Sprintf("%d", sumYes), nil
}
