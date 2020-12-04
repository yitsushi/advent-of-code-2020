package day04

import (
	"bufio"
	"io"
	"strings"
)

// Solver is the main solver.
type Solver struct {
	input []passport
}

// SetInput receives the input and parses its content.
func (d *Solver) SetInput(input io.Reader) error {
	scanner := bufio.NewScanner(input)

	passportData := []string{}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			d.input = append(d.input, parsePassport(passportData))
			passportData = []string{}

			continue
		}

		passportData = append(passportData, strings.Split(line, " ")...)
	}

	if len(passportData) > 0 {
		d.input = append(d.input, parsePassport(passportData))
	}

	return scanner.Err()
}
