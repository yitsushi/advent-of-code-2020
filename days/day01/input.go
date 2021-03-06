package day01

import (
	"bufio"
	"io"
	"strconv"
)

// SetInput receives the input and parses its content.
func (d *Solver) SetInput(input io.Reader) error {
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		value, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			return err
		}

		d.input = append(d.input, value)
	}

	return scanner.Err()
}
