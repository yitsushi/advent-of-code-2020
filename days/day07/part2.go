package day07

import (
	"fmt"

	"github.com/yitsushi/advent-of-code-2020/pkg/perf"
)

// Part2 for this day.
func (d *Solver) Part2() (string, error) {
	defer perf.Duration(perf.Track("Part2"))

	count := d.fetch(d.target)

	// Subtract one as we don't count the shiny gold one.
	count -= 1

	return fmt.Sprintf("%d", count), nil
}

func (d *Solver) fetch(target string) uint64 {
	if !d.inventory.Has(target) || len(d.inventory.Get(target)) < 1 {
		return 1
	}

	sum := uint64(1)

	for _, item := range d.inventory.Get(target) {
		sum += item.Count * d.fetch(item.Bag.Name)
	}

	return sum
}
