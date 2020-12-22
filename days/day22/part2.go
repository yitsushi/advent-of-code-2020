package day22

import (
	"fmt"
	"sort"

	"github.com/sirupsen/logrus"
	"github.com/yitsushi/advent-of-code-2020/pkg/perf"
	"github.com/yitsushi/advent-of-code-2020/pkg/slice"
)

// Part2 for this day.
func (d *Solver) Part2() (string, error) {
	defer perf.Duration(perf.Track("Part2"))
	// defer profile.Start().Stop()

	winner := d.recursiveRound(d.players...)

	logrus.Infof("Player %d won.", winner.ID)

	return fmt.Sprintf("%d", winner.Score()), nil
}

func (d *Solver) recursiveRound(players ...*Player) *Player {
	visited := map[int][]string{}

	for _, player := range players {
		visited[player.ID] = []string{}
	}

	for {
		table := []drawnCard{}
		recurse := true

		same := true

		for _, player := range players {
			if _, found := slice.ContainsString(visited[player.ID], player.Deck.Hash()); !found {
				same = false

				break
			}
		}

		if same {
			return players[0]
		}

		for _, player := range players {
			visited[player.ID] = append(visited[player.ID], player.Deck.Hash())
			card, _ := player.Draw()

			if card.Value > player.Size() {
				recurse = false
			}

			table = append(table, drawnCard{Player: player, Card: card})
		}

		var winner *Player

		if recurse {
			winner = d.startRecursiveGame(table, players...)
		} else {
			sort.Slice(table, func(i, j int) bool { return table[i].Card.Value > table[j].Card.Value })

			winner = table[0].Player
		}

		winnerCard, rest := slitTable(winner.ID, table)

		winner.Stash(winnerCard)
		winner.Stash(rest...)

		for _, player := range players {
			if player.Size() == 0 {
				return winner
			}
		}
	}
}

func (d *Solver) startRecursiveGame(table []drawnCard, players ...*Player) *Player {
	newPlayers := []*Player{}

	for idx, player := range players {
		newPlayers = append(newPlayers, player.Clone(table[idx].Card.Value))
	}

	subWinner := d.recursiveRound(newPlayers...)

	for _, player := range players {
		if player.ID == subWinner.ID {
			return player
		}
	}

	return nil
}

func slitTable(winnerID int, table []drawnCard) (Card, []Card) {
	var (
		winnerCard Card
		rest       = []Card{}
	)

	for _, card := range table {
		if card.Player.ID == winnerID {
			winnerCard = card.Card

			continue
		}

		rest = append(rest, card.Card)
	}

	return winnerCard, rest
}
