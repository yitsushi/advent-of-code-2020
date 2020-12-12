package day12

import "fmt"

// UnknownInstruction occurs when the input contains an invalid command.
type UnknownInstruction struct {
	Input string
}

func (e UnknownInstruction) Error() string {
	return fmt.Sprintf("Unknown instruction: %s", e.Input)
}
