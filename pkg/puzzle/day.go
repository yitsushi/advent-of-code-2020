package puzzle

import "io"

// Day is the interface for daily solutions.
type Day interface {
	Part1() (string, error)
	Part2() (string, error)
	SetInput(value io.Reader) error
}
