package day07

import (
	"bufio"
	"io"
)

// Solver is the main solver.
type Solver struct {
	rules     []Rule
	inventory inventory
	target    string
}

// SetInput receives the input and parses its content.
func (d *Solver) SetInput(input io.Reader) error {
	scanner := bufio.NewScanner(input)

	d.inventory = newInventory()

	for scanner.Scan() {
		rule := Matcher(scanner.Text())

		if !d.inventory.Has(rule.Bag.Name) {
			d.inventory.Store(rule.Bag.Name, rule.Contains)
		}

		d.rules = append(d.rules, rule)
	}

	d.target = "shiny gold"

	return scanner.Err()
}
