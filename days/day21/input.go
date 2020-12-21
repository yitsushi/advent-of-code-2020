package day21

import (
	"bufio"
	"io"
	"regexp"
	"strings"
)

// Solver is the main solver.
type Solver struct {
	bank Bank
}

// SetInput receives the input and parses its content.
func (d *Solver) SetInput(input io.Reader) error {
	scanner := bufio.NewScanner(input)
	macther := regexp.MustCompile(`^(.*) \(contains (.*)\)`)
	d.bank = NewBank()

	for scanner.Scan() {
		text := scanner.Text()

		food := NewFood()

		matches := macther.FindStringSubmatch(text)

		for _, ingredient := range strings.Split(matches[1], " ") {
			food.AddIngredients(d.bank.GetIngredient(ingredient))
		}

		food.AddAllergens(strings.Split(matches[2], ", ")...)

		d.bank.AddFood(&food)
	}

	return scanner.Err()
}
