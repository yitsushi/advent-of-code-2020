package day19

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/yitsushi/advent-of-code-2020/pkg/bullshit"
	"github.com/yitsushi/advent-of-code-2020/pkg/perf"
)

// Part1 for this day.
func (d *Solver) Part1() (string, error) {
	defer perf.Duration(perf.Track("Part1"))

	return fmt.Sprintf("%d", d.CountValidStrings()), nil
}

// Part2 for this day.
func (d *Solver) Part2() (string, error) {
	defer perf.Duration(perf.Track("Part2"))

	d.ruleGroups[8] = newRuleGroupFromString("8: 42 | 42 8")
	d.ruleGroups[11] = newRuleGroupFromString("11: 42 31 | 42 11 31")

	return fmt.Sprintf("%d", d.CountValidStrings()), nil
}

func (d *Solver) CountValidStrings() int64 {
	var counter int64

	matcher := d.GenerateRexep()

	for _, input := range d.input {
		if matcher.MatchString(input) {
			counter++
		}
	}

	return counter
}

func (d *Solver) GenerateRexep() *regexp.Regexp {
	keepGoing := true
	cycles := 0
	matcher := regexp.MustCompile(`\d+`)
	remover := regexp.MustCompile(`\(\(([ab])\)\)`)

	for id, group := range d.ruleGroups {
		group.Regexp = group.String()
		keepGoing = true

		d.ruleGroups[id] = group
	}

	maxDepth := map[int64]int{
		8:  10,
		11: 10,
	}

	for keepGoing && cycles < 100 {
		keepGoing = false
		group := d.ruleGroups[0]
		result := matcher.ReplaceAllStringFunc(group.Regexp, func(match string) string {
			id := bullshit.DropErrorInt64(strconv.ParseInt(match, 10, 64))

			switch id {
			case 8:
				maxDepth[8]--
				if maxDepth[8] < 0 {
					return match
				}
			case 11:
				maxDepth[11]--
				if maxDepth[11] < 0 {
					return match
				}
			}

			return fmt.Sprintf("(%s)", d.ruleGroups[id].Regexp)
		})
		result = remover.ReplaceAllString(result, "($1)")

		if result != group.Regexp {
			keepGoing = true
			group.Regexp = result
			d.ruleGroups[0] = group
		}

		cycles++
	}

	logrus.Infof("Done in %d rounds.", cycles)

	exp := strings.Replace(d.ruleGroups[0].Regexp, " ", "", -1)

	return regexp.MustCompile(fmt.Sprintf("^%s$", exp))
}
