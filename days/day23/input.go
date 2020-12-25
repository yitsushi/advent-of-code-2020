package day23

import (
	"bytes"
	"io"
	"io/ioutil"
)

// Solver is the main solver.
type Solver struct {
	cups []int
}

const (
	asciiInt = 48
)

// SetInput receives the input and parses its content.
func (d *Solver) SetInput(input io.Reader) error {
	text, _ := ioutil.ReadAll(input)

	text = bytes.TrimSpace(text)

	for _, ch := range text {
		value := ch - asciiInt
		d.cups = append(d.cups, int(value))
	}

	return nil
}
