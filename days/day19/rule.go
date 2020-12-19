package day19

import (
	"strconv"
	"strings"

	"github.com/yitsushi/advent-of-code-2020/pkg/bullshit"
	"github.com/yitsushi/advent-of-code-2020/pkg/parsing"
)

// RuleGroup is one rule group. One line in the input is one group.
type RuleGroup struct {
	ID       int64
	RuleSets []RuleSet
	Resolved []string
}

// RuleSet is a set if Rules.
type RuleSet struct {
	Rules   []Rule
	RefList []int64
}

// Rule is a single byte or referense to other RuleGroup.
type Rule struct {
	Value byte
	Ref   int64
}

func newRuleGroupFromString(text string) RuleGroup {
	group := RuleGroup{
		ID:       -1,
		RuleSets: []RuleSet{},
		Resolved: []string{},
	}

	parts := strings.SplitN(text, ": ", 2)
	group.ID = bullshit.DropErrorInt64(strconv.ParseInt(parts[0], 10, 64))

	for _, set := range strings.Split(parts[1], " | ") {
		group.RuleSets = append(group.RuleSets, newRuleSetFromString(set))
	}

	return group
}

func newRuleSetFromString(text string) RuleSet {
	set := RuleSet{
		Rules: []Rule{},
	}

	if strings.Contains(text, "\"") {
		set.Rules = append(set.Rules, Rule{Value: text[1], Ref: -1})

		return set
	}

	references, _ := parsing.Int64Slice(text, " ")

	for _, ref := range references {
		set.Rules = append(set.Rules, Rule{Ref: ref})
		set.RefList = append(set.RefList, ref)
	}

	return set
}
