package day11

import (
	"bufio"
	"io"
)

// Solver is the main solver.
type Solver struct {
	ferry Ferry
}

// SetInput receives the input and parses its content.
func (d *Solver) SetInput(input io.Reader) error {
	scanner := bufio.NewScanner(input)
	d.ferry = NewFerry()

	for scanner.Scan() {
		row, err := ParseRow(scanner.Text())
		if err != nil {
			return err
		}

		d.ferry.AddRow(row)
	}

	return scanner.Err()
}
