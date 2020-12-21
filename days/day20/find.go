package day20

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/yitsushi/advent-of-code-2020/pkg/slice"
)

const (
	hasTobeMiddle = 0
	hasTobeSide   = 1
	hasTobeCorner = 2

	rotateTwo = 2
)

func (d *Solver) findPostionGroups() {
	for _, target := range d.tiles {
		allPossible := []string{}

		for _, tile := range d.tiles {
			if tile.ID == target.ID {
				continue
			}

			allPossible = append(allPossible, tile.AllSidesAllVariant()...)
		}

		noMatch := 0

		for _, block := range target.AllInBlocks() {
			found := false

			for _, side := range block {
				if _, match := slice.ContainsString(allPossible, side); match {
					found = true

					break
				}
			}

			if !found {
				noMatch++
			}
		}

		switch noMatch {
		case hasTobeMiddle:
			d.image.AddMiddle(target)
		case hasTobeSide:
			d.image.AddSide(target)
		case hasTobeCorner:
			d.image.AddCorner(target)
		default:
			panic(fmt.Sprintf("To many missing pieces: %d", noMatch))
		}
	}
}

// FindTopLeftCorner tries to find the top left tile of the image.
func (d *Solver) FindTopLeftCorner() {
	for _, target := range d.image.Corner {
		allPossible := []string{}

		for _, tile := range d.image.Side {
			allPossible = append(allPossible, tile.AllSidesAllVariant()...)
		}

		if _, foundBottom := slice.ContainsString(allPossible, target.BottomRow()); foundBottom {
			if _, found := slice.ContainsString(allPossible, target.RightColumn()); found {
				d.image.Write(target)

				return
			}
		}

		if _, foundTop := slice.ContainsString(allPossible, target.TopRow()); foundTop {
			if _, found := slice.ContainsString(allPossible, target.LeftColumn()); found {
				target.Rotate(rotateTwo)
				d.image.Write(target)

				return
			}
		}
	}

	logrus.Panic("Couldn't find Top Left Corner")
}

// FindRow tries to fill a row with tiles.
func (d *Solver) FindRow() bool {
	if len(d.image.Data[d.image.Head]) == 0 {
		return d.FindNextBottom()
	}

	for len(d.image.Data[d.image.Head]) != d.image.GridSize {
		if !d.FindNextRight() {
			return false
		}
	}

	d.image.Head++

	return d.image.Head < d.image.GridSize
}

// FindNextRight tries to find the next tile that can fit netx to the previous tile.
func (d *Solver) FindNextRight() bool {
	target := d.image.Data[d.image.Head][len(d.image.Data[d.image.Head])-1]

	for _, tile := range d.tiles {
		if tile.ID == target.ID || tile.Placed {
			continue
		}

		if fit, instructions := target.CanFitRight(tile); fit {
			instructions.Execute(tile)
			d.image.Write(tile)

			return true
		}
	}

	return false
}

// FindNextBottom tries to find the next tile that can fit the bottom of the first
// tile in the previous row.
func (d *Solver) FindNextBottom() bool {
	target := d.image.Data[d.image.Head-1][0]

	for _, tile := range d.tiles {
		if tile.ID == target.ID || tile.Placed {
			continue
		}

		if fit, instructions := target.CanFitBottom(tile); fit {
			instructions.Execute(tile)
			d.image.Write(tile)

			return true
		}
	}

	return false
}
