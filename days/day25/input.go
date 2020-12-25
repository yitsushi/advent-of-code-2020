package day25

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
		text := scanner.Text()
		key := bullshit.DropErrorInt64(strconv.ParseInt(text, 10, 64))
		d.input = append(d.input, key)
	}

	return scanner.Err()
}
