package day20

import (
	"strings"
)

// Image is an image.
type Image struct {
	Corner   []*Tile
	Side     []*Tile
	Middle   []*Tile
	GridSize int
	Data     [][]*Tile
	Head     int
}

// NewImage creates a new Image.
func NewImage() *Image {
	return &Image{
		Corner:   []*Tile{},
		Side:     []*Tile{},
		Middle:   []*Tile{},
		GridSize: 0,
	}
}

// AddCorner saves a corner piece.
func (image *Image) AddCorner(tile *Tile) {
	image.Corner = append(image.Corner, tile)
}

// AddSide saves a side piece.
func (image *Image) AddSide(tile *Tile) {
	image.Side = append(image.Side, tile)
}

// AddMiddle saves a middle piece.
func (image *Image) AddMiddle(tile *Tile) {
	image.Middle = append(image.Middle, tile)
}

// SetGridSize to set the grid size (N x N).
func (image *Image) SetGridSize(size int) {
	image.GridSize = size
	image.Data = make([][]*Tile, size)
}

// Write a tile to head.
func (image *Image) Write(tile *Tile) {
	tile.Placed = true
	image.Data[image.Head] = append(image.Data[image.Head], tile)
}

// Merge image tiles into one tile.
func (image *Image) Merge() *Tile {
	full := NewTile(-1)

	for _, row := range image.Data {
		for idx := 1; idx < len(row[0].Data[0])-1; idx++ {
			fullLine := []string{}

			for _, tile := range row {
				fullLine = append(fullLine, tile.Data[idx][1:len(tile.Data[idx])-1])
			}

			full.Data = append(full.Data, strings.Join(fullLine, ""))
		}
	}

	return full
}
