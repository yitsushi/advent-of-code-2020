package day19

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/yitsushi/advent-of-code-2020/pkg/bullshit"
)

const (
	maxRecursion = 10
	maxCycles    = 100
	track8       = 8
	track11      = 11
)

func (d *Solver) generateRexep() *regexp.Regexp {
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
		8:  maxRecursion,
		11: maxRecursion,
	}

	for keepGoing && cycles < maxCycles {
		keepGoing = false
		group := d.ruleGroups[0]
		result := matcher.ReplaceAllStringFunc(group.Regexp, func(match string) string {
			id := bullshit.DropErrorInt64(strconv.ParseInt(match, 10, 64))

			switch id {
			case track8:
				maxDepth[track8]--
				if maxDepth[track8] < 0 {
					return match
				}
			case track11:
				maxDepth[track11]--
				if maxDepth[track11] < 0 {
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

	exp := strings.ReplaceAll(d.ruleGroups[0].Regexp, " ", "")

	return regexp.MustCompile(fmt.Sprintf("^%s$", exp))
}
