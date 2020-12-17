package day17

import (
	"github.com/yitsushi/advent-of-code-2020/pkg/math"
)

// Space is a 3D space where nodes can live.
type Space struct {
	storage       map[string]*Node
	neighborCache map[string][]*Node

	minValues math.Vector3D
	maxValues math.Vector3D
}

// NewSpace creates a new Space.
func NewSpace() Space {
	return Space{
		storage:       map[string]*Node{},
		neighborCache: map[string][]*Node{},
	}
}

// Nodes returns all known nodes.
func (s *Space) Nodes() []*Node {
	nodes := []*Node{}

	for _, node := range s.storage {
		nodes = append(nodes, node)
	}

	return nodes
}

// SelectNodes with given Active state.
func (s *Space) SelectNodes(active bool) []*Node {
	nodes := []*Node{}

	for _, node := range s.storage {
		if node.Active == active {
			nodes = append(nodes, node)
		}
	}

	return nodes
}

// Finalize all nodes.
func (s *Space) Finalize() {
	for _, node := range s.storage {
		node.Finalize()
	}
}

// Lookup a Node at a given coordinate.
func (s *Space) Lookup(coordinate math.Vector3D) *Node {
	return s.storage[coordinate.Hash()]
}

// Inspect a coordinate.
// If there is a Node, return with the Node.
// If no Node living there, create one.
func (s *Space) Inspect(coordinate math.Vector3D) *Node {
	node := s.Lookup(coordinate)
	if node == nil {
		node = &Node{
			Active:     false,
			Coordinate: coordinate,
		}

		s.minValues.MinimizeFrom(coordinate)
		s.maxValues.MaximizeFrom(coordinate)

		s.storage[coordinate.Hash()] = node
	}

	return node
}

// Neighbors for a given Node.
func (s *Space) Neighbors(node *Node) []*Node {
	nodes := []*Node{}

	if node == nil {
		return nodes
	}

	if cache, found := s.neighborCache[node.Coordinate.Hash()]; found {
		return cache
	}

	for _, neighbor := range node.Coordinate.Neighbors() {
		n := s.Inspect(neighbor)

		if n != nil {
			nodes = append(nodes, n)
		}
	}

	s.neighborCache[node.Coordinate.Hash()] = nodes

	return nodes
}

// SelectNeighbors with given Active state.
func (s *Space) SelectNeighbors(node *Node, active bool) []*Node {
	nodes := []*Node{}

	if node == nil {
		return nodes
	}

	for _, neighbor := range s.Neighbors(node) {
		if neighbor.Active == active {
			nodes = append(nodes, neighbor)
		}
	}

	return nodes
}
