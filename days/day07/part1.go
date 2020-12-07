package day07

import (
	"fmt"

	"github.com/yitsushi/advent-of-code-2020/pkg/perf"
)

// Part1 for this day.
func (d *Solver) Part1() (string, error) {
	defer perf.Duration(perf.Track("Part1"))

	all := [][]Bag{}

	alreadyHave := []string{}

	for _, rule := range d.rules {
		out := d.checkBag(rule.Bag)

		for _, path := range out {
			if len(path) <= 1 {
				continue
			}

			found := false

			for _, have := range alreadyHave {
				if path[0].Name == have {
					found = true

					break
				}
			}

			if !found {
				alreadyHave = append(alreadyHave, path[0].Name)
				all = append(all, path)
			}
		}
	}

	return fmt.Sprintf("%d", len(all)), nil
}

func (d *Solver) checkBag(bag Bag) [][]Bag {
	if bag.Name == d.target {
		return [][]Bag{{bag}}
	}

	bags := [][]Bag{}

	d.inventory.Get(bag.Name)

	current := []Bag{bag}

	for _, content := range d.inventory.Get(bag.Name) {
		var rest [][]Bag

		rest = d.checkBag(content.Bag)

		for _, rest := range rest {
			bags = append(bags, append(current, rest...))
		}
	}

	return bags
}
