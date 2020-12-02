package day02

import (
	"bufio"
	"fmt"
	"io"
)

// SetInput receives the input and parses its content.
func (d *Solver) SetInput(input io.Reader) error {
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()

		var (
			min, max int
			ch       byte
			pw       []byte
		)

		fmt.Sscanf(line, "%d-%d %c: %s", &min, &max, &ch, &pw)

		d.input = append(d.input, entry{
			Min:       min,
			Max:       max,
			Character: ch,
			Password:  pw,
		})
	}

	return scanner.Err()
}
