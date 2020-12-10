package day10

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/yitsushi/advent-of-code-2020/pkg/math"
	"github.com/yitsushi/advent-of-code-2020/pkg/perf"
)

const (
	adapterOffset = 3
	shortHop      = 1
	longHop       = 3
)

// Part1 for this day.
func (d *Solver) Part1() (string, error) {
	defer perf.Duration(perf.Track("Part1"))

	d.input = append(d.input, math.MaximumInt64(d.input)+adapterOffset)
	numberOfOneHop, numberOfThreeHop := 0, 0
	current := int64(0)

	for _, jolt := range d.input {
		switch jolt - current {
		case shortHop:
			numberOfOneHop++
		case longHop:
			numberOfThreeHop++
		}

		current = jolt
	}

	logrus.Debugf("number of one hops: %d", numberOfOneHop)
	logrus.Debugf("number of three hops: %d", numberOfThreeHop)

	return fmt.Sprintf("%d", numberOfOneHop*numberOfThreeHop), nil
}

// Part2 for this day.
func (d *Solver) Part2() (string, error) {
	defer perf.Duration(perf.Track("Part2"))

	pathCache := newCache()

	logrus.Info(d.input)

	target := math.MaximumInt64(d.input) + adapterOffset
	d.input = append(d.input, target)
	d.input = append([]int64{0}, d.input...)
	paths := walk(d.input, &pathCache, target)

	return fmt.Sprintf("%d", paths), nil
}
