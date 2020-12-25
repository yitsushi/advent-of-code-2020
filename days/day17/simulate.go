package day17

import "github.com/yitsushi/advent-of-code-2020/pkg/generic"

const (
	i2aValue = 3
	a2iMin   = 2
	a2iMax   = 3
)

func (d *Solver) simulate() {
	for cycle := 0; cycle < 6; cycle++ {
		queue := generic.NewQueue()
		changeQueue := generic.NewQueue()

		for _, node := range d.space.SelectNodes(true) {
			queue.Push(node)

			for _, n := range d.space.Neighbors(node) {
				queue.Push(n)
			}
		}

		for !queue.Empty() {
			node := queue.Pull().(*Node)

			activeNum := len(d.space.SelectNeighbors(node, true))
			activate := (!node.Active && activeNum == i2aValue)
			deactivate := (node.Active && activeNum != a2iMin && activeNum != a2iMax)

			if activate || deactivate {
				node.ToggleActive()
				changeQueue.Push(node)
			}
		}

		for !changeQueue.Empty() {
			node := changeQueue.Pull().(*Node)

			node.Finalize()
		}
	}
}
