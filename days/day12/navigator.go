package day12

import (
	"github.com/yitsushi/advent-of-code-2020/pkg/math"
)

// Instruction contains one instruction command with its value.
type Instruction struct {
	Command byte
	Value   float64
}

// Navigator is the navigation system.
type Navigator struct {
	Instructions []Instruction
	Ship         math.OrintedObject2D
}

// NewNavigator creates a new Navigator.
func NewNavigator(position, facing math.Vector2D) Navigator {
	return Navigator{
		Instructions: []Instruction{},
		Ship: math.OrintedObject2D{
			Position: position,
			Facing:   facing,
		},
	}
}

// AddInstruction saves an instruction at the end of the instruction list.
func (n *Navigator) AddInstruction(instruction Instruction) {
	n.Instructions = append(n.Instructions, instruction)
}

// Simulate the instruction list.
func (n *Navigator) Simulate(waypoint bool) {
	var target *math.Vector2D

	if waypoint {
		target = &n.Ship.Facing
	} else {
		target = &n.Ship.Position
	}

	for _, instruction := range n.Instructions {
		switch instruction.Command {
		case 'N':
			target.Y += instruction.Value
		case 'S':
			target.Y -= instruction.Value
		case 'E':
			target.X += instruction.Value
		case 'W':
			target.X -= instruction.Value
		case 'F':
			n.Ship.Move(instruction.Value)
		case 'L':
			n.Ship.Rotate(instruction.Value)
		case 'R':
			n.Ship.Rotate(-instruction.Value)
		}
	}
}
