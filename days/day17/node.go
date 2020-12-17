package day17

import "github.com/yitsushi/advent-of-code-2020/pkg/math"

type Node struct {
	Active     bool
	Coordinate math.Vector3D
	NextState  bool
}

func (n *Node) InstantSet(active bool) {
	n.Active = active
	n.NextState = active
}

func (n *Node) Activate() {
	n.NextState = true
}

func (n *Node) Dectivate() {
	n.NextState = false
}

func (n *Node) ToggleActive() {
	n.NextState = !n.Active
}

func (n *Node) Finalize() {
	n.Active = n.NextState
}
