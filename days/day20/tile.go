package day20

import (
	"github.com/yitsushi/advent-of-code-2020/pkg/slice"
)

// Tile is a tile.
type Tile struct {
	ID     int
	Data   []string
	Placed bool

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

// FlipX flips the tile on the X axes.
func (t *Tile) FlipX() {
	result := []string{}

	for idx := len(t.Data) - 1; idx >= 0; idx-- {
		result = append(result, t.Data[idx])
	}

	t.Data = result
}

// FlipY flips the tile on the Y axes.
func (t *Tile) FlipY() {
	for idx := 0; idx < len(t.Data); idx++ {
		t.Data[idx] = slice.ReverseString(t.Data[idx])
	}
}

// Rotate CCW.
func (t *Tile) Rotate(times int) {
	times %= 4

	for ; times > 0; times-- {
		result := make([]string, len(t.Data))

		for _, line := range t.Data {
			for idx := 0; idx < len(line); idx++ {
				result[len(line)-idx-1] += string(line[idx])
			}
		}

		t.Data = result
	}
}

// CanFitRight returns with true if given tile can fit the right side of the tile
// and if it can fit, what operations we have to do on the tile.
func (t *Tile) CanFitRight(tile *Tile) (bool, FitInstruction) {
	left, right := tile.LeftColumn(), tile.RightColumn()
	top, bottom := tile.TopRow(), tile.BottomRow()

	switch t.RightColumn() {
	case left:
		return true, FitInstruction{}
	case slice.ReverseString(left):
		return true, FitInstruction{FlipX: true}
	case right:
		return true, FitInstruction{FlipY: true}
	case slice.ReverseString(right):
		return true, FitInstruction{FlipX: true, FlipY: true}
	case top:
		return true, FitInstruction{Rotation: 1, FlipX: true}
	case slice.ReverseString(top):
		return true, FitInstruction{Rotation: 1}
	case bottom:
		return true, FitInstruction{Rotation: 1, FlipY: true, FlipX: true}
	case slice.ReverseString(bottom):
		return true, FitInstruction{Rotation: 1, FlipY: true}
	}

	return false, FitInstruction{}
}

// CanFitBottom returns with true if given tile can fit below the tile
// and if it can fit, what operations we have to do on the tile.
func (t *Tile) CanFitBottom(tile *Tile) (bool, FitInstruction) {
	left, right := tile.LeftColumn(), tile.RightColumn()
	top, bottom := tile.TopRow(), tile.BottomRow()

	switch t.BottomRow() {
	case left:
		return true, FitInstruction{Rotation: 1, FlipX: true}
	case slice.ReverseString(left):
		return true, FitInstruction{Rotation: 1, FlipY: true, FlipX: true}
	case right:
		return true, FitInstruction{Rotation: 1}
	case slice.ReverseString(right):
		return true, FitInstruction{Rotation: 1, FlipY: true}
	case top:
		return true, FitInstruction{}
	case slice.ReverseString(top):
		return true, FitInstruction{FlipY: true}
	case bottom:
		return true, FitInstruction{FlipX: true}
	case slice.ReverseString(bottom):
		return true, FitInstruction{FlipX: true, FlipY: true}
	}

	return false, FitInstruction{}
}

// FindMonsters on a tile.
func (t *Tile) FindMonsters() int {
	monsterCount := 0

	for idx := 1; idx < len(t.Data)-1; idx++ {
		line := t.Data[idx]

		for lidx := range line[:len(line)-len(monsterMiddle)] {
			nope := false

			for midx := range monsterMiddle {
				rp1c, rp1 := monsterTop[midx] == '#', monsterTop[midx]
				rp2c, rp2 := monsterMiddle[midx] == '#', monsterMiddle[midx]
				rp3c, rp3 := monsterBottom[midx] == '#', monsterBottom[midx]
				rs1 := t.Data[idx-1][lidx+midx]
				rs2 := t.Data[idx][lidx+midx]
				rs3 := t.Data[idx+1][lidx+midx]

				if (rp1c && rp1 != rs1) || (rp2c && rp2 != rs2) || (rp3c && rp3 != rs3) {
					nope = true

					break
				}
			}

			if !nope {
				monsterCount++
			}
		}
	}

	return monsterCount
}
