package day22

import (
	"fmt"
	"sort"

	"github.com/sirupsen/logrus"
	"github.com/yitsushi/advent-of-code-2020/pkg/perf"
)

type drawnCard struct {
	Player *Player
	Card   Card
}

// Part1 for this day.
func (d *Solver) Part1() (string, error) {
	defer perf.Duration(perf.Track("Part1"))

	winner := d.combat(d.players...)
	logrus.Infof("Player %d won.", winner.ID)

	return fmt.Sprintf("%d", winner.Score()), nil
}

func (d *Solver) combat(players ...*Player) *Player {
	for {
		table := []drawnCard{}

		for _, player := range players {
			card, _ := player.Draw()

			table = append(table, drawnCard{Player: player, Card: card})
		}

		sort.Slice(table, func(i, j int) bool { return table[i].Card.Value > table[j].Card.Value })

		logrus.Infof("Player %d won with %d", table[0].Player.ID, table[0].Card.Value)

		winner := table[0].Player

		for _, card := range table {
			winner.Stash(card.Card)
		}

		for _, player := range players {
			if player.Size() == 0 {
				return winner
			}
		}
	}
}
