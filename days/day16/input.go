package day16

import (
	"bufio"
	"io"
	"strings"

	"github.com/yitsushi/advent-of-code-2020/pkg/parsing"
)

const (
	criteriaSection = "ticket criteria list"
	myTicketSection = "my ticket"
	nearbySection   = "nearby tickets"
)

// Solver is the main solver.
type Solver struct {
	criteriaList  []*Criteria
	myTicket      Ticket
	nearbyTickets []Ticket
}

// SetInput receives the input and parses its content.
func (d *Solver) SetInput(input io.Reader) error {
	scanner := bufio.NewScanner(input)

	section := criteriaSection

	for scanner.Scan() {
		text := scanner.Text()

		switch text {
		case "your ticket:":
			section = myTicketSection
		case "nearby tickets:":
			section = nearbySection
		case "":
			continue
		default:
			err := d.parseLine(section, text)
			if err != nil {
				return err
			}
		}
	}

	return scanner.Err()
}

func (d *Solver) parseLine(section, text string) error {
	if section == criteriaSection {
		criteria := Criteria{Ranges: []Range{}}

		mainParts := strings.SplitN(text, ": ", 2)
		criteria.Name = mainParts[0]

		for _, c := range strings.Split(mainParts[1], " or ") {
			from, to, err := parsing.Int64Range(c, "-")
			if err != nil {
				return err
			}

			r := Range{From: from, To: to}
			criteria.Ranges = append(criteria.Ranges, r)
		}

		d.criteriaList = append(d.criteriaList, &criteria)

		return nil
	}

	list, err := parsing.Int64Slice(text, ",")
	if err != nil {
		return err
	}

	ticket := Ticket{Fields: list}

	if section == myTicketSection {
		d.myTicket = ticket
	}

	if section == nearbySection {
		d.nearbyTickets = append(d.nearbyTickets, ticket)
	}

	return nil
}
