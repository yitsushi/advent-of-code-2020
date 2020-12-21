package day21

import (
	"fmt"
	"sort"
	"strings"

	"github.com/yitsushi/advent-of-code-2020/pkg/perf"
)

// Part1 for this day.
func (d *Solver) Part1() (string, error) {
	defer perf.Duration(perf.Track("Part1"))

	d.bank.Finalize()

	counter := 0

	for _, ingredient := range d.bank.IngredientsWithoutAllergen() {
		counter += len(d.bank.AllFoodWithIngredient(ingredient))
	}

	return fmt.Sprintf("%d", counter), nil
}

// Part2 for this day.
func (d *Solver) Part2() (string, error) {
	defer perf.Duration(perf.Track("Part2"))

	d.bank.Finalize()

	sort.Strings(d.bank.Allergen)

	list := []string{}

	for _, allergen := range d.bank.Allergen {
		list = append(list, d.bank.IngredientWithAllegen(allergen))
	}

	return strings.Join(list, ","), nil
}
