package day22

// Player with ID and a Deck.
type Player struct {
	ID   int
	Deck Deck
}

// Deck with cards.
type Deck []Card

// Card with value.
type Card struct {
	Value int
}

// Draw a card.
func (p *Player) Draw() (Card, bool) {
	if len(p.Deck) < 1 {
		return Card{}, false
	}

	card, rest := p.Deck[0], p.Deck[1:]

	p.Deck = rest

	return card, true
}

// Stash new cards at the bottom of the deck.
func (p *Player) Stash(cards ...Card) {
	p.Deck = append(p.Deck, cards...)
}

// Size of the deck.
func (p *Player) Size() int {
	return len(p.Deck)
}

// Score of the player.
func (p *Player) Score() int {
	score := 0

	for idx := 1; idx <= len(p.Deck); idx++ {
		score += p.Deck[len(p.Deck)-idx].Value * idx
	}

	return score
}

// Clone player with sub-deck.
func (p *Player) Clone(size int) *Player {
	newDeck := make([]Card, size)
	copy(newDeck, p.Deck)

	return &Player{
		ID:   p.ID,
		Deck: newDeck,
	}
}

// Hash the deck.
func (d Deck) Hash() string {
	values := []byte{}
	for _, card := range d {
		values = append(values, byte(card.Value))
	}

	return string(values)
}
