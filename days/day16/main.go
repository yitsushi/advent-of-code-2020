package day16

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/yitsushi/advent-of-code-2020/pkg/perf"
)

// Part1 for this day.
func (d *Solver) Part1() (string, error) {
	defer perf.Duration(perf.Track("Part1"))

	invalidFields := []int64{}

	for _, ticket := range d.nearbyTickets {
		invalidFields = append(invalidFields, ticket.Validate(d.criteriaList)...)
	}

	sum := int64(0)

	for _, v := range invalidFields {
		sum += v
	}

	return fmt.Sprintf("%d", sum), nil
}

// Part2 for this day.
func (d *Solver) Part2() (string, error) {
	defer perf.Duration(perf.Track("Part2"))

	departureFields := []string{}

	for _, c := range d.criteriaList {
		if strings.HasPrefix(c.Name, "departure") {
			departureFields = append(departureFields, c.Name)
		}
	}

	trackers := []*Tracker{}

	for _, ticket := range append(d.nearbyTickets, d.myTicket) {
		if len(ticket.Validate(d.criteriaList)) == 0 {
			for idx, field := range ticket.Fields {
				if len(trackers) <= idx {
					trackers = append(trackers, NewTracker(idx))
				}

				trackers[idx].AddValue(field)
			}
		}
	}

	for !allFieldTracked(trackers, departureFields) {
		for _, tracker := range trackers {
			candidates := tracker.Candidates(d.criteriaList)

			if len(candidates) == 1 {
				candidates[0].Taken = true
				tracker.Name = candidates[0].Name

				logrus.Debugf("%d has to be %s", tracker.Index, tracker.Name)

				continue
			}

			logrus.Debugf("%d has %d candidates", tracker.Index, len(candidates))
		}
	}

	solution := int64(1)

	for _, t := range trackers {
		if t.Name != "" {
			logrus.Infof("%d is %s", t.Index, t.Name)
		} else {
			logrus.Infof("%d is unknown", t.Index)
		}
	}

	for _, field := range departureFields {
		for _, t := range trackers {
			if t.Name == field {
				solution *= t.OnMyTicket()
			}
		}
	}

	return fmt.Sprintf("%d", solution), nil
}

func allFieldTracked(trackers []*Tracker, fields []string) bool {
	for _, f := range fields {
		found := false

		for _, t := range trackers {
			if t.Name == f {
				found = true

				break
			}
		}

		if !found {
			return false
		}
	}

	return true
}
