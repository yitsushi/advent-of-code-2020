package day05

import (
	"bufio"
	"io"
)

// Solver is the main solver.
type Solver struct {
	input []Ticket
}

// SetInput receives the input and parses its content.
func (d *Solver) SetInput(input io.Reader) error {
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		ticket := NewTicket(scanner.Text())

		if err := ticket.Validate(); err != nil {
			return err
		}

		d.input = append(d.input, ticket)
	}

	return scanner.Err()
}
