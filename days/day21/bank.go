package day21

import (
	"github.com/yitsushi/advent-of-code-2020/pkg/slice"
)

// Bank is the food bank with all the information.
type Bank struct {
	Ingredient map[string]*Ingredient
	Food       []*Food
	Allergen   []string
}

// NewBank creates a new Bank.
func NewBank() Bank {
	return Bank{
		Ingredient: map[string]*Ingredient{},
		Food:       []*Food{},
		Allergen:   []string{},
	}
}

// GetIngredient finds an ingredient in the bank, or registers a new one.
func (b *Bank) GetIngredient(name string) *Ingredient {
	if ingredient, found := b.Ingredient[name]; found {
		return ingredient
	}

	ingredient := NewIngredient(name)

	b.Ingredient[name] = &ingredient

	return b.Ingredient[name]
}

// AddFood registers a food list in the bank.
func (b *Bank) AddFood(food *Food) {
	b.Food = append(b.Food, food)

	b.AddAllergens(food.Allergens...)
}

// AddAllergens registers an allergen list in the bank.
func (b *Bank) AddAllergens(allergens ...string) {
	for _, allergen := range allergens {
		if _, found := slice.ContainsString(b.Allergen, allergen); found {
			continue
		}

		b.Allergen = append(b.Allergen, allergen)
	}
}

// AllFoodWithAllegen returns with all foods with given allergen.
func (b *Bank) AllFoodWithAllegen(allergen string) []*Food {
	foodList := []*Food{}

	for _, food := range b.Food {
		if _, found := slice.ContainsString(food.Allergens, allergen); found {
			foodList = append(foodList, food)
		}
	}

	return foodList
}

// AllFoodWithIngredient returns with all foods with given ingredient.
func (b *Bank) AllFoodWithIngredient(ingredient string) []*Food {
	foodList := []*Food{}

	for _, food := range b.Food {
		if _, found := slice.ContainsString(food.IngredientList(), ingredient); found {
			foodList = append(foodList, food)
		}
	}

	return foodList
}

// IngredientsWithoutAllergen returns ingredients from the bank
// without registered allergen.
func (b *Bank) IngredientsWithoutAllergen() []string {
	list := []string{}

	for name, ingredient := range b.Ingredient {
		if ingredient.Allergen == "" {
			list = append(list, name)
		}
	}

	return list
}

// Finalize the bank and try to guess ingredients and their allergens.
func (b *Bank) Finalize() {
	for {
		found := 0

		for _, allergen := range b.Allergen {
			ingredientLists := [][]string{}

			for _, food := range b.AllFoodWithAllegen(allergen) {
				ingredientLists = append(ingredientLists, food.IngredientList())
			}

			common := slice.IntersectString(ingredientLists...)
			target := slice.IntersectString(common, b.IngredientsWithoutAllergen())

			if len(target) == 1 {
				b.GetIngredient(target[0]).Allergen = allergen

				found++

				continue
			}
		}

		if found == 0 {
			break
		}
	}
}

// IngredientWithAllegen returns with the name of the ingredient
// for a given allergen.
func (b *Bank) IngredientWithAllegen(name string) string {
	result := ""

	for _, ingredient := range b.Ingredient {
		if ingredient.Allergen == name {
			result = ingredient.Name

			break
		}
	}

	return result
}
