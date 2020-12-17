package day17

import (
	"fmt"

	"github.com/yitsushi/advent-of-code-2020/pkg/math"
	"github.com/yitsushi/advent-of-code-2020/pkg/perf"
)

// Part1 for this day.
func (d *Solver) Part1() (string, error) {
	defer perf.Duration(perf.Track("Part1"))

	d.simulate()

	return fmt.Sprintf("%d", len(d.space.SelectNodes(true))), nil
}

// Part2 for this day.
func (d *Solver) Part2() (string, error) {
	defer perf.Duration(perf.Track("Part2"))

	newSpace := NewSpace()

	for _, node := range d.space.Nodes() {
		values := node.Coordinate.Values()

		newNode := newSpace.Inspect(math.Vector4D{
			X: values[0],
			Y: values[1],
			Z: values[2],
			W: 0,
		})

		newNode.InstantSet(node.Active)
	}

	d.space = newSpace

	d.simulate()

	return fmt.Sprintf("%d", len(d.space.SelectNodes(true))), nil
}
