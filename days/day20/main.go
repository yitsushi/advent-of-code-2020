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

	for _, row := range d.image.Data {
		chain := []string{}

		for _, tile := range row {
			chain = append(chain, fmt.Sprintf("%d", tile.ID))
		}

		logrus.Info(chain)
	}

	fullImage := d.image.Merge()

	for _, line := range fullImage.Data {
		logrus.Debug(line)
	}

	numberOfMonsters := 0

	for numberOfMonsters == 0 {
		fullImage.Rotate(1)

		numberOfMonsters = len(fullImage.FindMonsters())
		if numberOfMonsters > 0 {
			break
		}

		fullImage.FlipX()

		numberOfMonsters = len(fullImage.FindMonsters())
		if numberOfMonsters > 0 {
			break
		}

		fullImage.FlipX()
		fullImage.FlipY()

		numberOfMonsters = len(fullImage.FindMonsters())
		if numberOfMonsters > 0 {
			break
		}

		fullImage.FlipY()
	}

	counter := -numberOfMonsters * monsterLength

	for _, line := range fullImage.Data {
		counter += strings.Count(line, "#")
	}

	return fmt.Sprintf("%d", counter), nil
}
