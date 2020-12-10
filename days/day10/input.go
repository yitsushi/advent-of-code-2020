package day10

import (
	"bufio"
	"io"
	"sort"
	"strconv"
)

// Solver is the main solver.
type Solver struct {
	input []int64
}

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

	sort.SliceStable(d.input, func(i, j int) bool {
		return d.input[i] < d.input[j]
	})

	return scanner.Err()
}
