package day20

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/yitsushi/advent-of-code-2020/pkg/perf"
	"github.com/yitsushi/advent-of-code-2020/pkg/puzzle"
)

const (
	numberOfCorners = 4
	maxRotation     = 4
)

// Part1 for this day.
func (d *Solver) Part1() (string, error) {
	defer perf.Duration(perf.Track("Part1"))

	d.findPostionGroups()

	logrus.Infof("The image is %dx%d", d.image.GridSize, d.image.GridSize)
	logrus.Infof("Corner tiles has %d tiles", len(d.image.Corner))

	if len(d.image.Corner) != numberOfCorners {
		return "", puzzle.NoSolution{}
	}

	product := int64(1)

	for _, tile := range d.image.Corner {
		product *= int64(tile.ID)
	}

	return fmt.Sprintf("%d", product), nil
}

// Part2 for this day.
func (d *Solver) Part2() (string, error) {
	defer perf.Duration(perf.Track("Part2"))

	d.findPostionGroups()
	d.FindTopLeftCorner()

	for d.FindRow() {
	}

	fullImage := d.image.Merge()

	for _, line := range fullImage.Data {
		logrus.Debug(line)
	}

	highest := 0
	rotations := 0

	for {
		fullImage.Rotate(1)
		rotations++

		if rotations > maxRotation {
			break
		}

		if res := fullImage.FindMonsters(); highest < res {
			highest = res
		}

		ops := []func(){
			fullImage.FlipX,
			func() {
				fullImage.FlipX()
				fullImage.FlipY()
			},
			fullImage.FlipY,
		}

		for _, op := range ops {
			op()

			if res := fullImage.FindMonsters(); highest < res {
				highest = res
			}
		}
	}

	counter := -highest * monsterLength

	for _, line := range fullImage.Data {
		counter += strings.Count(line, "#")
	}

	return fmt.Sprintf("%d", counter), nil
}
