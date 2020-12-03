package day03

import (
	"fmt"

	"github.com/yitsushi/advent-of-code-2020/pkg/perf"
)

func (d *Solver) checkRoute(vec vector2D) (int, error) {
	treeCount := 0

	d.Map.Reset()

	for d.Map.Step(vec) {
		if d.Map.CurrentSquare() == treeSpace {
			treeCount++

			continue
		}

		if d.Map.CurrentSquare() != openSpace {
			return 0, UnknownSquare{Value: d.Map.CurrentSquare()}
		}
	}

	return treeCount, nil
}

// Part1 for this day.
func (d *Solver) Part1() (string, error) {
	defer perf.Duration(perf.Track("Part1"))

	//nolint:gomnd
	treeCount, err := d.checkRoute(vector2D{X: 3, Y: 1})

	return fmt.Sprintf("%d", treeCount), err
}

// Part2 for this day.
func (d *Solver) Part2() (string, error) {
	defer perf.Duration(perf.Track("Part2"))

	//nolint:gomnd
	possibleSlopes := []vector2D{
		{X: 1, Y: 1},
		{X: 3, Y: 1},
		{X: 5, Y: 1},
		{X: 7, Y: 1},
		{X: 1, Y: 2},
	}

	solution := 1

	for _, slope := range possibleSlopes {
		treeCount, err := d.checkRoute(slope)
		if err != nil {
			return "", err
		}

		solution *= treeCount
	}

	return fmt.Sprintf("%d", solution), nil
}
