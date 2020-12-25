package day24

import (
	"github.com/yitsushi/advent-of-code-2020/pkg/generic"
	"github.com/yitsushi/advent-of-code-2020/pkg/math"
)

const (
	flipCount = 2
)

// Grid is basically the floor with tiles.
type Grid struct {
	Tiles map[math.Vector3D]*Tile
}

// NewGrid creates a new grid.
func NewGrid() *Grid {
	grid := &Grid{
		Tiles: map[math.Vector3D]*Tile{},
	}

	grid.Get(math.Vector3D{X: 0, Y: 0, Z: 0})

	return grid
}

// OneDayLater calculates the tile pattern for the next day.
func (grid *Grid) OneDayLater() {
	newTiles := map[math.Vector3D]*Tile{}
	updateQueue := generic.NewQueue()

	for _, tile := range grid.All(Black) {
		updateQueue.Push(tile.Coordinate)

		for _, item := range tile.Neighbors() {
			updateQueue.Push(item)
		}
	}

	for !updateQueue.Empty() {
		tile := grid.Get(updateQueue.Pull().(math.Vector3D))
		black, _ := grid.NeighborColorCount(tile)

		if tile.Color {
			if black == 1 || black == flipCount {
				newTiles[tile.Coordinate] = NewTile(tile.Coordinate).Flip()
			}

			continue
		}

		if black == flipCount {
			tile.Flip()

			newTiles[tile.Coordinate] = NewTile(tile.Coordinate).Flip()
		}
	}

	grid.Tiles = newTiles

	for _, tile := range newTiles {
		tile.Finalize()
	}
}

// NeighborColorCount returns with the number of back and white
// neighbors.
func (grid *Grid) NeighborColorCount(tile *Tile) (int, int) {
	black, white := 0, 0

	for _, n := range tile.Neighbors() {
		if grid.Get(n).Color {
			black++
		} else {
			white++
		}
	}

	return black, white
}

// All tiles with color.
func (grid *Grid) All(color Color) []*Tile {
	tiles := []*Tile{}

	for _, tile := range grid.Tiles {
		if tile.Color == color {
			tiles = append(tiles, tile)
		}
	}

	return tiles
}

// Count tiles with color.
func (grid *Grid) Count(color Color) int {
	counter := 0

	for _, tile := range grid.Tiles {
		if tile.Color == color {
			counter++
		}
	}

	return counter
}

// FlipFromInstruction executes all instructions for one tile.
func (grid *Grid) FlipFromInstruction(directions []Direction) {
	target := grid.Base()

	for _, direction := range directions {
		target = target.NeightborTo(direction)
	}

	current := grid.Get(target.Coordinate)

	current.Flip().Finalize()
}

// Get a file on a given coordinate.
func (grid *Grid) Get(coordinate math.Vector3D) *Tile {
	if tile, found := grid.Tiles[coordinate]; found {
		return tile
	}

	grid.Tiles[coordinate] = NewTile(coordinate)

	return grid.Tiles[coordinate]
}

// Base tile.
func (grid *Grid) Base() *Tile {
	return grid.Get(math.Vector3D{X: 0, Y: 0, Z: 0})
}
