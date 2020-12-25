package day24

import (
	"github.com/yitsushi/advent-of-code-2020/pkg/math"
)

// Tile colors.
const (
	Black = true
	White = false
)

// Color of the tile.
type Color bool

// Tile is a tile.
type Tile struct {
	Coordinate math.Vector3D
	Color      Color
	NextColor  Color
}

// NewTile creates a new tile.
func NewTile(coordinate math.Vector3D) *Tile {
	return &Tile{
		Color:      false,
		Coordinate: coordinate,
	}
}

// NeightborTo given direction.
func (t *Tile) NeightborTo(direction Direction) *Tile {
	switch direction {
	case East:
		return NewTile(t.East())
	case West:
		return NewTile(t.West())
	case NorthEast:
		return NewTile(t.NorthEast())
	case NorthWest:
		return NewTile(t.NorthWest())
	case SouthEast:
		return NewTile(t.SouthEast())
	case SouthWest:
		return NewTile(t.SouthWest())
	}

	return t
}

// Flip tile.
func (t *Tile) Flip() *Tile {
	t.NextColor = !t.Color

	return t
}

// Finalize changes.
func (t *Tile) Finalize() {
	if t.Color != t.NextColor {
		t.Color = t.NextColor
	}
}

// East neighbors.
func (t *Tile) East() math.Vector3D {
	return math.Vector3D{
		X: t.Coordinate.X - 1,
		Y: t.Coordinate.Y + 1,
		Z: t.Coordinate.Z,
	}
}

// West neighbors.
func (t *Tile) West() math.Vector3D {
	return math.Vector3D{
		X: t.Coordinate.X + 1,
		Y: t.Coordinate.Y - 1,
		Z: t.Coordinate.Z,
	}
}

// NorthEast neighbors.
func (t *Tile) NorthEast() math.Vector3D {
	return math.Vector3D{
		X: t.Coordinate.X,
		Y: t.Coordinate.Y + 1,
		Z: t.Coordinate.Z - 1,
	}
}

// NorthWest neighbors.
func (t *Tile) NorthWest() math.Vector3D {
	return math.Vector3D{
		X: t.Coordinate.X + 1,
		Y: t.Coordinate.Y,
		Z: t.Coordinate.Z - 1,
	}
}

// SouthWest neighbors.
func (t *Tile) SouthWest() math.Vector3D {
	return math.Vector3D{
		X: t.Coordinate.X,
		Y: t.Coordinate.Y - 1,
		Z: t.Coordinate.Z + 1,
	}
}

// SouthEast neighbors.
func (t *Tile) SouthEast() math.Vector3D {
	return math.Vector3D{
		X: t.Coordinate.X - 1,
		Y: t.Coordinate.Y,
		Z: t.Coordinate.Z + 1,
	}
}

// Neighbors of the tile.
func (t *Tile) Neighbors() []math.Vector3D {
	return []math.Vector3D{
		t.East(),
		t.West(),
		t.NorthEast(),
		t.NorthWest(),
		t.SouthEast(),
		t.SouthWest(),
	}
}
