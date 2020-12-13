package day13

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/yitsushi/advent-of-code-2020/pkg/bullshit"
)

// Solver is the main solver.
type Solver struct {
	earliest int64
	busList  []int64
}

// SetInput receives the input and parses its content.
func (d *Solver) SetInput(input io.Reader) error {
	scanner := bufio.NewScanner(input)

	if scanner.Scan() {
		d.earliest = bullshit.DropErrorInt64(strconv.ParseInt(scanner.Text(), 10, 64))
	}

	if scanner.Scan() {
		for _, value := range strings.Split(scanner.Text(), ",") {
			id, err := strconv.ParseInt(value, 10, 16)
			if err != nil {
				id = -1
			}

			d.busList = append(d.busList, id)
		}
	}

	return scanner.Err()
}
