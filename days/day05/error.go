package day05

import "fmt"

const (
	// InvalidTicketFormat is the template format for invalid ticket errors.
	InvalidTicketFormat = "Invalid Ticket %s; Position %d has '%c' character"
)

// InvalidTicket happens when a Ticket Serial contains at least one invalid character.
type InvalidTicket struct {
	Serial    string
	Position  int
	Character rune
}

func (e InvalidTicket) Error() string {
	return fmt.Sprintf("Invalid Ticket %s; Position %d has '%c' character", e.Serial, e.Position, e.Character)
}
