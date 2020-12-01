package day01

import (
	"io"
	"strconv"
	"strings"
)

// SetInput receives the input and parses its content.
func (d *Solver) SetInput(input io.Reader) error {
	buf := new(strings.Builder)

	_, err := io.Copy(buf, input)
	if err != nil {
		return err
	}

	d.input = []int64{}

	for _, line := range strings.Split(buf.String(), "\n") {
		value, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			return err
		}

		d.input = append(d.input, value)
	}

	return nil
}
