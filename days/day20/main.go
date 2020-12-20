package day20

import (
	"fmt"

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

	logrus.Infof("The image is %dx%d", d.image.NumberOfRows, d.image.NumberOfRows)
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

	return "", puzzle.NotImplemented{}
}
