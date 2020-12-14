package puzzle

import "fmt"

// UnkownPartError occurs when given part is not defined.
type UnkownPartError struct {
	Day  Day
	Part int
}

func (e UnkownPartError) Error() string {
	return fmt.Sprintf("unknown part: %T :: %d", e.Day, e.Part)
}

// NoInputError occurs when we have no input values.
type NoInputError struct{}

func (e NoInputError) Error() string {
	return "no input, no puzzle"
}

// NoSolution occurs when we are unable to find a solution.
type NoSolution struct{}

func (e NoSolution) Error() string {
	return "2020404 - solution not found"
}

// NotImplemented occurs when something is not implemented.
type NotImplemented struct{}

func (e NotImplemented) Error() string {
	return "2020501 - not implemented"
}
