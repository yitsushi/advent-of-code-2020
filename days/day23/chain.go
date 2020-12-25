package day23

import (
	"github.com/yitsushi/advent-of-code-2020/pkg/math"
)

const (
	cupsInHand = 3
)

// Chain of cups.
type Chain struct {
	Head *Cup
	Max  int

	cache map[int]*Cup
}

// NewChainOf n cups.
func NewChainOf(cups []int, numberOfItems int) *Chain {
	head := &Cup{Value: cups[0]}
	last := head
	chain := &Chain{
		Head:  head,
		Max:   numberOfItems,
		cache: map[int]*Cup{},
	}

	for _, value := range cups[1:] {
		new := &Cup{Value: value}
		last.Next = new
		last = new
		chain.cache[value] = new
	}

	for value := math.MaximumInt(cups) + 1; value <= numberOfItems; value++ {
		new := &Cup{Value: value}
		last.Next = new
		last = new
		chain.cache[value] = new
	}

	last.Next = head

	return chain
}

// Nth cup from Head.
func (chain *Chain) Nth(n int) *Cup {
	target := chain.Head

	for ; n > 0; n-- {
		target = target.Next
	}

	return target
}

// Find cup with label.
func (chain *Chain) Find(n int) *Cup {
	if cup, found := chain.cache[n]; found {
		return cup
	}

	target := chain.Head.Next

	for {
		if target.Value == n {
			return target
		}

		if target == chain.Head {
			return nil
		}

		target = target.Next
	}
}

// AllFrom constructs a slice of int values starting from
// a given cup.
func (chain *Chain) AllFrom(n int) []int {
	cup := chain.Find(n)
	list := []int{cup.Value}
	current := cup.Next

	for {
		if current.Value == cup.Value {
			break
		}

		list = append(list, current.Value)
		current = current.Next
	}

	return list
}

func (chain *Chain) crabShuffle() {
	hand := chain.Head.Next
	chain.Head.Next = chain.Nth(cupsInHand + 1)
	hand.Nth(cupsInHand - 1).Next = hand

	var (
		target      *Cup
		targetValue = chain.Head.Value
	)

	for {
		targetValue--

		if targetValue < 1 {
			targetValue = chain.Max
		}

		if found := hand.Find(targetValue); found != nil {
			continue
		}

		target = chain.Find(targetValue)

		break
	}

	hand.Nth(cupsInHand - 1).Next = target.Next
	target.Next = hand

	chain.Head = chain.Head.Next
}
