package day03

import "fmt"

// UnknownSquare occurs when we try to step onto an unknown square on the map.
type UnknownSquare struct {
	Value byte
}

func (e UnknownSquare) Error() string {
	return fmt.Sprintf("Unknown square: %c", e.Value)
}
