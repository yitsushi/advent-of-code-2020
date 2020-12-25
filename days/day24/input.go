package day24

import (
	"bufio"
	"io"
)

// Solver is the main solver.
type Solver struct {
	instructions [][]Direction
}

// SetInput receives the input and parses its content.
func (d *Solver) SetInput(input io.Reader) error {
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		text := scanner.Text()

		// text = strings.Replace(text, "swne", "", -1)
		// text = strings.Replace(text, "nesw", "", -1)
		// text = strings.Replace(text, "senw", "", -1)
		// text = strings.Replace(text, "senw", "", -1)

		instructions := []Direction{}

		for idx := 0; idx < len(text); {
			current := text[idx : idx+1]

			if idx+1 < len(text) {
				current += string(text[idx+1])
			}

			instruction, offset, err := directionFromString(current)
			if err != nil {
				return err
			}

			instructions = append(instructions, instruction)
			idx += offset
		}

		d.instructions = append(d.instructions, instructions)
	}

	return scanner.Err()
}
