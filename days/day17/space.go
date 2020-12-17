package day17

import (
	"github.com/sirupsen/logrus"
	"github.com/yitsushi/advent-of-code-2020/pkg/math"
)

type Space struct {
	storage       map[string]*Node
	neighborCache map[string][]*Node

	minValues math.Vector3D
	maxValues math.Vector3D
}

func NewSpace() Space {
	return Space{
		storage:       map[string]*Node{},
		neighborCache: map[string][]*Node{},
	}
}

func (s *Space) Nodes() []*Node {
	nodes := []*Node{}

	for _, node := range s.storage {
		nodes = append(nodes, node)
	}

	return nodes
}

func (s *Space) SelectNodes(active bool) []*Node {
	nodes := []*Node{}

	for _, node := range s.storage {
		if node.Active == active {
			nodes = append(nodes, node)
		}
	}

	return nodes
}

func (s *Space) Finalize() {
	for _, node := range s.storage {
		node.Finalize()
	}
}

func (s *Space) Lookup(coordinate math.Vector3D) *Node {
	return s.storage[coordinate.Hash()]
}

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

func (s *Space) Neighbors(node *Node) []*Node {
	nodes := []*Node{}

	if node == nil {
		return nodes
	}

	// if cache, found := s.neighborCache[node.Coordinate.Hash()]; found {
	// 	return cache
	// }

	for _, neighbor := range node.Coordinate.Neighbors() {
		n := s.Inspect(neighbor)

		if n != nil {
			nodes = append(nodes, n)
		}
	}

	// s.neighborCache[node.Coordinate.Hash()] = nodes

	return nodes
}

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

func (s *Space) Display() [][]string {
	logrus.Debug(s.minValues)
	logrus.Debug(s.maxValues)

	layers := [][]string{}

	for z := s.minValues.Z; z <= s.maxValues.Z; z++ {
		lines := []string{}

		for y := s.minValues.Y; y <= s.maxValues.Y; y++ {
			line := []rune{}
			for x := s.minValues.X; x <= s.maxValues.X; x++ {
				ch := inactiveStateCharacter

				if s.Inspect(math.Vector3D{X: x, Y: y, Z: z}).Active {
					ch = activeStateCharacter
				}

				line = append(line, ch)
			}

			lines = append(lines, string(line))
		}

		layers = append(layers, lines)
	}

	return layers
}

func (s *Space) Extend() {
	for _, node := range s.storage {
		for _, n := range node.Coordinate.Neighbors() {
			s.Inspect(n)
		}
	}
}
