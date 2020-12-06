package day06

import (
	"bytes"

	"github.com/sirupsen/logrus"
)

// Group represents a group from the plane.
type Group struct {
	YesAnswers       [][]byte
	UniqueYesAnswers []byte
	UnanimousYes     []byte
}

// NewGroup creates a new group.
func NewGroup() Group {
	return Group{
		YesAnswers:       [][]byte{},
		UniqueYesAnswers: []byte{},
		UnanimousYes:     []byte{},
	}
}

// Feed the answer of one passenger.
func (g *Group) Feed(answers []byte) {
	g.YesAnswers = append(g.YesAnswers, answers)

	for _, ch := range answers {
		if !bytes.Contains(g.UniqueYesAnswers, []byte{ch}) {
			g.UniqueYesAnswers = append(g.UniqueYesAnswers, ch)
		}
	}
}

// Finalize the results.
func (g *Group) Finalize() {
	full := bytes.Join(g.YesAnswers, []byte{' '})
	numberOfMembers := len(g.YesAnswers)

	for _, answer := range g.UniqueYesAnswers {
		if bytes.Count(full, []byte{answer}) == numberOfMembers {
			g.UnanimousYes = append(g.UnanimousYes, answer)
		}
	}

	logrus.Debugf(
		"Unanimous: %s (members: %d)",
		g.UnanimousYes, numberOfMembers)
}
