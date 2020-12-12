package day12

import (
	"fmt"

	"github.com/yitsushi/advent-of-code-2020/pkg/math"
	"github.com/yitsushi/advent-of-code-2020/pkg/perf"
)

const (
	waypointX = 10
	waypointY = 1
)

// Part1 for this day.
func (d *Solver) Part1() (string, error) {
	defer perf.Duration(perf.Track("Part1"))

	d.navigator.Simulate(false)

	return fmt.Sprintf("%.0f", d.navigator.Ship.Position.Manhattan()), nil
}

// Part2 for this day.
func (d *Solver) Part2() (string, error) {
	defer perf.Duration(perf.Track("Part2"))

	d.navigator.Ship.Facing = math.Vector2D{X: waypointX, Y: waypointY}
	d.navigator.Simulate(true)

	return fmt.Sprintf("%.0f", d.navigator.Ship.Position.Manhattan()), nil
}
