package day09

import (
	"bufio"
	"io"
	"strconv"

	"github.com/yitsushi/advent-of-code-2020/pkg/bullshit"
)

// Solver is the main solver.
type Solver struct {
	input []int64
}

// SetInput receives the input and parses its content.
func (d *Solver) SetInput(input io.Reader) error {
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		d.input = append(d.input, bullshit.DropErrorInt64(strconv.ParseInt(scanner.Text(), 10, 64)))
	}

	return scanner.Err()
}
