package day17

import (
	"fmt"

	"github.com/yitsushi/advent-of-code-2020/pkg/perf"
)

// Part1 for this day.
func (d *Solver) Part1() (string, error) {
	defer perf.Duration(perf.Track("Part1"))

	for cycle := 0; cycle < 6; cycle++ {
		queue := NewUpdateQueue()
		changeQueue := NewUpdateQueue()

		for _, node := range d.space3D.SelectNodes(true) {
			queue.Push(node)
			queue.Push(d.space3D.Neighbors(node)...)
		}

		for !queue.Empty() {
			node := queue.Pull()

			activeNum := len(d.space3D.SelectNeighbors(node, true))
			activate := (!node.Active && activeNum == i2aValue)
			deactivate := (node.Active && activeNum != a2iMin && activeNum != a2iMax)

			if activate || deactivate {
				node.ToggleActive()
				changeQueue.Push(node)
			}
		}

		for !changeQueue.Empty() {
			node := changeQueue.Pull()

			node.Finalize()
		}
	}

	countActive := 0

	for _, node := range d.space3D.Nodes() {
		if node.Active {
			countActive++
		}
	}

	return fmt.Sprintf("%d", countActive), nil
}
