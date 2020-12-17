package day17

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/yitsushi/advent-of-code-2020/pkg/perf"
	"github.com/yitsushi/advent-of-code-2020/pkg/puzzle"
)

const (
	i2aValue = 3
	a2iMin   = 2
	a2iMax   = 3
)

// Part1 for this day.
func (d *Solver) Part1() (string, error) {
	defer perf.Duration(perf.Track("Part1"))

	for cycle := 0; cycle < 6; cycle++ {
		queue := NewUpdateQueue()
		changeQueue := NewUpdateQueue()

		for _, node := range d.space.SelectNodes(true) {
			queue.Push(node)
			queue.Push(d.space.Neighbors(node)...)
		}

		for !queue.Empty() {
			node := queue.Pull()

			activeNum := len(d.space.SelectNeighbors(node, true))
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

	for _, node := range d.space.Nodes() {
		if node.Active {
			countActive++
		}
	}

	return fmt.Sprintf("%d", countActive), nil
}

// Part2 for this day.
func (d *Solver) Part2() (string, error) {
	defer perf.Duration(perf.Track("Part2"))

	return "", puzzle.NotImplemented{}
}

func (d *Solver) draw() {
	for idx, layer := range d.space.Display() {
		logrus.Debugf("Layer #%d", idx)
		for _, line := range layer {
			logrus.Debug(line)
		}
		logrus.Debug("")
	}
}
