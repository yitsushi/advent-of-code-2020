package day19

import (
	"fmt"

	"github.com/yitsushi/advent-of-code-2020/pkg/perf"
)

// Part1 for this day.
func (d *Solver) Part1() (string, error) {
	defer perf.Duration(perf.Track("Part1"))

	return fmt.Sprintf("%d", d.countValidStrings()), nil
}

// Part2 for this day.
func (d *Solver) Part2() (string, error) {
	defer perf.Duration(perf.Track("Part2"))

	d.ruleGroups[8] = newRuleGroupFromString("8: 42 | 42 8")
	d.ruleGroups[11] = newRuleGroupFromString("11: 42 31 | 42 11 31")

	return fmt.Sprintf("%d", d.countValidStrings()), nil
}

func (d *Solver) countValidStrings() int64 {
	var counter int64

	matcher := d.generateRexep()

	for _, input := range d.input {
		if matcher.MatchString(input) {
			counter++
		}
	}

	return counter
}
