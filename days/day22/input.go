package day22

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// Solver is the main solver.
type Solver struct {
	players []*Player
}

// SetInput receives the input and parses its content.
func (d *Solver) SetInput(input io.Reader) error {
	scanner := bufio.NewScanner(input)

	d.players = []*Player{}

	for scanner.Scan() {
		text := scanner.Text()

		if strings.HasPrefix(text, "Player") {
			player := &Player{Deck: []Card{}}

			fmt.Sscanf(text, "Player %d:", &player.ID)

			d.players = append(d.players, player)

			continue
		}

		value, err := strconv.Atoi(text)
		if err != nil {
			continue
		}

		d.players[len(d.players)-1].Deck = append(
			d.players[len(d.players)-1].Deck,
			Card{Value: value},
		)
	}

	return scanner.Err()
}
