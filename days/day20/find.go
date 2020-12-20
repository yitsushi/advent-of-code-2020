package day20

import (
	"fmt"

	"github.com/yitsushi/advent-of-code-2020/pkg/slice"
)

const (
	hasTobeMiddle = 0
	hasTobeSide   = 1
	hasTobeCorner = 2
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
