package day07

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/yitsushi/advent-of-code-2020/pkg/bullshit"
)

// Matcher parses a sentence into a Rule.
func Matcher(sentence string) Rule {
	mainMatcher := regexp.MustCompile(`([a-z]+) ([a-z]+) bags contain`)
	subMatcher := regexp.MustCompile(`(\d+) ([a-z]+) ([a-z]+) bag`)

	mainResult := mainMatcher.FindStringSubmatch(sentence)
	subResults := subMatcher.FindAllStringSubmatch(sentence, -1)

	rule := Rule{
		Bag: Bag{
			Tone:  mainResult[1],
			Color: mainResult[2],
			Name:  fmt.Sprintf("%s %s", mainResult[1], mainResult[2]),
		},
		Contains: []ContentStatement{},
	}

	for _, match := range subResults {
		rule.Contains = append(rule.Contains, ContentStatement{
			Count: bullshit.DropErrorUint64(strconv.ParseUint(match[1], 10, 64)),
			Bag: Bag{
				Tone:  match[2],
				Color: match[3],
				Name:  fmt.Sprintf("%s %s", match[2], match[3]),
			},
		})
	}

	return rule
}
