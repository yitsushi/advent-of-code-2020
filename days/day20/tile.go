package day20

import "github.com/yitsushi/advent-of-code-2020/pkg/slice"

// Tile is a tile.
type Tile struct {
	ID     int
	Data   []string
	Top    *Tile
	Left   *Tile
	Bottom *Tile
	Right  *Tile

	all []string
}

// NewTile creates a new Tile.
func NewTile(id int) *Tile {
	return &Tile{
		ID:   id,
		Data: []string{},
		all:  []string{},
	}
}

// Finalize a tile.
// Reason: cache.
func (t *Tile) Finalize() {
	t.all = append(t.all, t.AllTop()...)
	t.all = append(t.all, t.AllBottom()...)
	t.all = append(t.all, t.AllLeft()...)
	t.all = append(t.all, t.AllRight()...)
}

// TopRow of the tile.
func (t *Tile) TopRow() string {
	return t.Data[0]
}

// BottomRow of the tile.
func (t *Tile) BottomRow() string {
	return t.Data[len(t.Data)-1]
}

// LeftColumn of the tile.
func (t *Tile) LeftColumn() string {
	column := []byte{}

	for _, b := range t.Data {
		column = append(column, b[0])
	}

	return string(column)
}

// RightColumn of the tile.
func (t *Tile) RightColumn() string {
	column := []byte{}

	for _, b := range t.Data {
		column = append(column, b[len(b)-1])
	}

	return string(column)
}

// AllSidesAllVariant of the tile.
func (t *Tile) AllSidesAllVariant() []string {
	return t.all
}

// AllTop of the tile.
func (t *Tile) AllTop() []string {
	sides := []string{}

	sides = append(sides, t.TopRow())
	sides = append(sides, slice.ReverseString(t.TopRow()))

	return sides
}

// AllBottom of the tile.
func (t *Tile) AllBottom() []string {
	sides := []string{}

	sides = append(sides, t.BottomRow())
	sides = append(sides, slice.ReverseString(t.BottomRow()))

	return sides
}

// AllLeft of the tile.
func (t *Tile) AllLeft() []string {
	sides := []string{}

	sides = append(sides, t.LeftColumn())
	sides = append(sides, slice.ReverseString(t.LeftColumn()))

	return sides
}

// AllRight of the tile.
func (t *Tile) AllRight() []string {
	sides := []string{}

	sides = append(sides, t.RightColumn())
	sides = append(sides, slice.ReverseString(t.RightColumn()))

	return sides
}

// AllInBlocks of the tile.
func (t *Tile) AllInBlocks() [][]string {
	result := [][]string{}

	result = append(result, t.AllTop())
	result = append(result, t.AllBottom())
	result = append(result, t.AllLeft())
	result = append(result, t.AllRight())

	return result
}
