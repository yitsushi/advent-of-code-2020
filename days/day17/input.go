package day17

import (
	"bufio"
	"io"

	"github.com/yitsushi/advent-of-code-2020/pkg/math"
)

// Solver is the main solver.
type Solver struct {
	space3D Space
}

const activeStateCharacter = '#'

// SetInput receives the input and parses its content.
func (d *Solver) SetInput(input io.Reader) error {
	scanner := bufio.NewScanner(input)

	d.space3D = NewSpace()
	y := float64(0)

	for scanner.Scan() {
		for x, ch := range scanner.Text() {
			node := d.space3D.Inspect(math.Vector3D{
				X: float64(x),
				Y: y,
				Z: 0,
			})

			node.InstantSet(ch == activeStateCharacter)
		}

		y++
	}

	return scanner.Err()
}
