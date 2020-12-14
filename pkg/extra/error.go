package extra

import "fmt"

// UnkownDayError occurs when given day is not defined.
type UnkownDayError struct {
	Day int
}

func (e UnkownDayError) Error() string {
	return fmt.Sprintf("Unknown day: %d", e.Day)
}
