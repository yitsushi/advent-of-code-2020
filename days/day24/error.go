package day24

import "fmt"

// UnknownDirection occurs when a given direction is unknown.
type UnknownDirection struct {
	Direction string
}

func (e UnknownDirection) Error() string {
	return fmt.Sprintf("unknown direction: %s", e.Direction)
}
