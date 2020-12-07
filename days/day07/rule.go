package day07

// Rule is one sentence parsed as a rule from regulations handbook.
type Rule struct {
	Bag      Bag
	Contains []ContentStatement
}
