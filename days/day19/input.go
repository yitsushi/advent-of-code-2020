package day19

import (
	"bufio"
	"io"
)

// Solver is the main solver.
type Solver struct {
	input      []string
	ruleGroups map[int64]RuleGroup
}

// SetInput receives the input and parses its content.
func (d *Solver) SetInput(input io.Reader) error {
	scanner := bufio.NewScanner(input)

	parseRules := true
	d.input = []string{}
	d.ruleGroups = map[int64]RuleGroup{}

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			parseRules = false

			continue
		}

		if !parseRules {
			d.input = append(d.input, scanner.Text())

			continue
		}

		group := newRuleGroupFromString(line)

		d.ruleGroups[group.ID] = group
	}

	return scanner.Err()
}
