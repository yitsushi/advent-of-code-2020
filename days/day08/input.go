package day08

import (
	"bufio"
	"io"
)

// Solver is the main solver.
type Solver struct {
	machine Machine
}

// SetInput receives the input and parses its content.
func (d *Solver) SetInput(input io.Reader) error {
	scanner := bufio.NewScanner(input)

	source := []Instruction{}

	for scanner.Scan() {
		ins, err := ParseInstruction(scanner.Text())
		if err != nil {
			return err
		}

		source = append(source, ins)
	}

	d.machine.Load(source)
	d.machine.Reset()

	return scanner.Err()
}
