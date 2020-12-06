package day06

import (
	"bufio"
	"io"
)

// Solver is the main solver.
type Solver struct {
	input []Group
}

// SetInput receives the input and parses its content.
func (d *Solver) SetInput(input io.Reader) error {
	scanner := bufio.NewScanner(input)

	group := NewGroup()

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			group.Finalize()

			d.input = append(d.input, group)
			group = NewGroup()

			continue
		}

		group.Feed([]byte(line))
	}

	if len(group.YesAnswers) > 0 {
		group.Finalize()

		d.input = append(d.input, group)
	}

	return scanner.Err()
}
