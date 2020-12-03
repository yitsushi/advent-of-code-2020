package day03

import (
	"bufio"
	"io"
)

const (
	openSpace = '.'
	treeSpace = '#'
)

// Solver is the main solver.
type Solver struct {
	Map snowField
}

// SetInput receives the input and parses its content.
func (d *Solver) SetInput(input io.Reader) error {
	scanner := bufio.NewScanner(input)

	area := []string{}

	for scanner.Scan() {
		area = append(area, scanner.Text())
	}

	d.Map = newMap(area)

	return scanner.Err()
}
