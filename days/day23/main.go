package day23

import (
	"strconv"

	"github.com/yitsushi/advent-of-code-2020/pkg/perf"
	"github.com/yitsushi/advent-of-code-2020/pkg/slice"
)

const (
	cupsPart1   = 9
	roundsPart1 = 100
	cupsPart2   = 1000000
	roundsPart2 = 10000000
)

// Part1 for this day.
func (d *Solver) Part1() (string, error) {
	defer perf.Duration(perf.Track("Part1"))

	chain := NewChainOf(d.cups, cupsPart1)

	for counter := 0; counter < roundsPart1; counter++ {
		chain.crabShuffle()
	}

	final := chain.AllFrom(1)

	return slice.JoinInt(final[1:], ""), nil
}

// Part2 for this day.
func (d *Solver) Part2() (string, error) {
	defer perf.Duration(perf.Track("Part2"))

	chain := NewChainOf(d.cups, cupsPart2)

	for counter := 0; counter < roundsPart2; counter++ {
		chain.crabShuffle()
	}

	target := chain.Find(1).Next

	return strconv.Itoa(target.Value * target.Next.Value), nil
}
