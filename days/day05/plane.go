package day05

import (
	"github.com/sirupsen/logrus"
)

const (
	maxNumberOfRows    = 127
	maxNumberOfColumns = 7

	// This is silly.
	two = 2
)

// Plane is a... plane.
type Plane struct {
	Tickets map[int]Ticket
	Rows    [][]bool
}

// NewPlane create a new Plane.
func NewPlane() Plane {
	rows := [][]bool{}

	for i := 0; i < 128; i++ {
		rows = append(rows, make([]bool, 8))
	}

	return Plane{
		Tickets: make(map[int]Ticket),
		Rows:    rows,
	}
}

type block struct {
	From int
	To   int
}

// BookSeat sets a seat booked for given ticket.
func (p *Plane) BookSeat(ticket Ticket) {
	section := block{From: 0, To: maxNumberOfRows}

	for _, ch := range ticket.Serial[:7] {
		midPoint := (section.From + section.To) / two

		if ch == Front {
			section.To = midPoint
		} else {
			section.From = midPoint + 1
		}
	}

	ticket.Row = section.From
	section = block{From: 0, To: maxNumberOfColumns}

	for _, ch := range ticket.Serial[7:] {
		midPoint := (section.From + section.To) / two

		if ch == Left {
			section.To = midPoint
		} else {
			section.From = midPoint + 1
		}
	}

	ticket.Column = section.From
	ticket.ID = calculateTicketID(ticket.Row, ticket.Column)

	logrus.Debugf("%+v", ticket)

	p.Rows[ticket.Row][ticket.Column] = true
	p.Tickets[ticket.ID] = ticket
}

// Draw renders the plane and its seats.
//   '*' -> booked
//   '_' -> free
func (p Plane) Draw() []string {
	output := []string{}

	for _, row := range p.Rows {
		block := ""

		for _, seat := range row[:4] {
			if seat {
				block += "*"
			} else {
				block += "_"
			}
		}

		block += " "

		for _, seat := range row[4:] {
			if seat {
				block += "*"
			} else {
				block += "_"
			}
		}

		output = append(output, block)
	}

	return output
}
