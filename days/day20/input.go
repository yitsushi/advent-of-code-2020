package day20

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strings"
)

// Solver is the main solver.
type Solver struct {
	tiles []*Tile
	image *Image
}

// SetInput receives the input and parses its content.
func (d *Solver) SetInput(input io.Reader) error {
	scanner := bufio.NewScanner(input)

	d.tiles = []*Tile{}
	d.image = NewImage()

	var tile *Tile

	for scanner.Scan() {
		text := scanner.Text()

		if strings.HasPrefix(text, "Tile") {
			id := 0

			fmt.Sscanf(text, "Tile %d:", &id)

			tile = NewTile(id)

			continue
		}

		if len(text) == 0 {
			tile.Finalize()

			d.tiles = append(d.tiles, tile)
			tile = nil

			continue
		}

		tile.Data = append(tile.Data, text)
	}

	if tile != nil {
		tile.Finalize()

		d.tiles = append(d.tiles, tile)
	}

	d.image.NumberOfRows = int(math.Sqrt(float64(len(d.tiles))))

	return scanner.Err()
}
