package day15

import (
	"fmt"

	"github.com/yitsushi/advent-of-code-2020/pkg/perf"
)

const (
	targetPart1 = 2020
	targetPart2 = 30000000
)

// Part1 for this day.
func (d *Solver) Part1() (string, error) {
	defer perf.Duration(perf.Track("Part1"))

	return fmt.Sprintf("%d", d.walk(targetPart1)), nil
}

// Part2 for this day.
func (d *Solver) Part2() (string, error) {
	defer perf.Duration(perf.Track("Part2"))

	return fmt.Sprintf("%d", d.walk(targetPart2)), nil
}

func (d *Solver) walk(to int) int64 {
	var (
		history    map[int64]int = map[int64]int{}
		lastNumber int64
	)

	for index, value := range d.input[:len(d.input)-1] {
		history[value] = index + 1
	}

	lastNumber = d.input[len(d.input)-1]

	for i := len(d.input); i < to; i++ {
		var newNumber int64

		if index, exists := history[lastNumber]; exists {
			newNumber = int64(i - index)
		}

		history[lastNumber] = i
		lastNumber = newNumber
	}

	return lastNumber
}
