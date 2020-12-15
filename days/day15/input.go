package day15

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

// Solver is the main solver.
type Solver struct {
	input []int64
}

// SetInput receives the input and parses its content.
func (d *Solver) SetInput(input io.Reader) error {
	scanner := bufio.NewScanner(input)

	if !scanner.Scan() {
		return scanner.Err()
	}

	for _, number := range strings.Split(scanner.Text(), ",") {
		value, err := strconv.ParseInt(number, 10, 64)
		if err != nil {
			return err
		}

		d.input = append(d.input, value)
	}

	return scanner.Err()
}
