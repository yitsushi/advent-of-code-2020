package day19

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/yitsushi/advent-of-code-2020/pkg/perf"
)

// Part1 for this day.
func (d *Solver) Part1() (string, error) {
	defer perf.Duration(perf.Track("Part1"))

	var counter int

	for _, input := range d.input {
		if d.isPossible(input, 0) {
			counter++
		}
	}

	return fmt.Sprintf("%d", counter), nil
}

// Part2 for this day.
func (d *Solver) Part2() (string, error) {
	defer perf.Duration(perf.Track("Part2"))

	d.ruleGroups[8] = newRuleGroupFromString("8: 42 | 42 8")
	d.ruleGroups[11] = newRuleGroupFromString("11: 42 31 | 42 11 31")

	var counter int

	for _, input := range d.input {
		if d.isPossible(input, 0) {
			logrus.Warn(input)
			counter++
		}
	}

	return fmt.Sprintf("%d", counter), nil
}

func (d *Solver) isPossible(pattern string, target int64) bool {
	logrus.Infof("Is it possible? %s", pattern)

	if pattern == "" {
		logrus.Println("Found")
	}

	s, ok := d.checkGroup(pattern, d.ruleGroups[target])

	return ok && len(pattern) == s
}

func (d *Solver) checkGroup(pattern string, group RuleGroup) (int, bool) {
	for _, set := range group.RuleSets {
		logrus.Debugf("Check Rule: %d => %v", group.ID, set.RefList)
		if shift, ok := d.checkRuleSet(pattern, set); ok {
			return shift, true
		}
	}

	return 0, false
}

func (d *Solver) checkRuleSet(pattern string, set RuleSet) (int, bool) {
	logrus.Debugf("Check RuleSet on %s: %v", pattern, set.RefList)

	if len(pattern) < 1 {
		logrus.Error("End of pattern")
	}

	if len(pattern) < 1 {
		return 0, false
	}

	shift := 0

	for _, rule := range set.Rules {
		logrus.Debug(pattern[shift:], rule)
		if rule.Ref == -1 {
			if rule.Value != byte(pattern[shift]) {
				return 0, false
			}

			shift++
			continue
		}

		s, check := d.checkGroup(pattern[shift:], d.ruleGroups[rule.Ref])
		if !check {
			return 0, false
		}

		shift += s
	}

	return shift, true
}
