package day05

import "strings"

// Sections.
const (
	Front = 'F'
	Back  = 'B'
	Left  = 'L'
	Right = 'R'
)

const (
	possibleValues        = "FRLB"
	ticketRowIDMultiplier = 8
)

// Ticket representsa one ticket.
type Ticket struct {
	ID     int
	Serial string
	Row    int
	Column int
}

// NewTicket creates as new ticket from a Serial.
func NewTicket(serial string) Ticket {
	return Ticket{
		Serial: serial,
		Row:    0,
		Column: 0,
	}
}

// Validate a ticket.
func (t Ticket) Validate() error {
	for idx, ch := range t.Serial {
		//nolint:gocritic
		if !strings.Contains(possibleValues, string(ch)) {
			return InvalidTicket{
				Serial:    t.Serial,
				Position:  idx,
				Character: ch,
			}
		}
	}

	return nil
}

func calculateTicketID(row, column int) int {
	return row*ticketRowIDMultiplier + column
}
