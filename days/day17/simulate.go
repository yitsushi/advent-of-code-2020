package day17

const (
	i2aValue = 3
	a2iMin   = 2
	a2iMax   = 3
)

func (d *Solver) simulate() {
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
}
