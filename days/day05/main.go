package day05

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/yitsushi/advent-of-code-2020/pkg/perf"
	"github.com/yitsushi/advent-of-code-2020/pkg/puzzle"
)

// Part1 for this day.
func (d *Solver) Part1() (string, error) {
	defer perf.Duration(perf.Track("Part1"))

	plane := d.checkInTickets()
	max := 0

	for _, ticket := range plane.Tickets {
		if max < ticket.ID {
			max = ticket.ID
		}
	}

	return fmt.Sprintf("%d", max), nil
}

// Part2 for this day.
func (d *Solver) Part2() (string, error) {
	defer perf.Duration(perf.Track("Part2"))

	plane := d.checkInTickets()
	picture := plane.Draw()

	logrus.Debugf("\n%s", strings.Join(picture, "\n"))

	for row, block := range picture {
		if strings.Count(block, "_") == 1 {
			column := strings.Index(block, "_")
			id := calculateTicketID(row, column)

			if _, prev := plane.Tickets[id-1]; prev {
				if _, next := plane.Tickets[id+1]; next {
					logrus.Debugf("%+v", plane.Tickets[id-1])
					logrus.Debugf("%+v", plane.Tickets[id+1])

					return fmt.Sprintf("%d", id), nil
				}
			}
		}
	}

	return "", puzzle.NoSolution{}
}

func (d *Solver) checkInTickets() Plane {
	plane := NewPlane()

	for _, ticket := range d.input {
		plane.BookSeat(ticket)
	}

	return plane
}
