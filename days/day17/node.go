package day17

import "github.com/yitsushi/advent-of-code-2020/pkg/math"

// Node is a single power cube.
type Node struct {
	Active     bool
	NextState  bool
	Coordinate math.Vector
}

// InstantSet sets the Active status without needing to finalize it.
func (n *Node) InstantSet(active bool) {
	n.Active = active
	n.NextState = active
}

// Activate the node.
// It's not instant, calling Finalize is required.
func (n *Node) Activate() {
	n.NextState = true
}

// Dectivate the node.
// It's not instant, calling Finalize is required.
func (n *Node) Dectivate() {
	n.NextState = false
}

// ToggleActive toggles the state of the node.
// It's not instant, calling Finalize is required.
func (n *Node) ToggleActive() {
	n.NextState = !n.Active
}

// Finalize changes.
func (n *Node) Finalize() {
	n.Active = n.NextState
}
