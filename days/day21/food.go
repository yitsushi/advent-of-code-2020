package day21

// Food is a food.
type Food struct {
	Ingredients []*Ingredient
	Allergens   []string
}

// NewFood create a new Food.
func NewFood() Food {
	return Food{
		Ingredients: []*Ingredient{},
		Allergens:   []string{},
	}
}

// AddIngredients for a food.
func (f *Food) AddIngredients(ingredients ...*Ingredient) {
	f.Ingredients = append(f.Ingredients, ingredients...)
}

// AddAllergens for a food.
func (f *Food) AddAllergens(allergens ...string) {
	f.Allergens = append(f.Allergens, allergens...)
}

// IngredientList for a food.
func (f *Food) IngredientList() []string {
	list := []string{}

	for _, ingredient := range f.Ingredients {
		list = append(list, ingredient.Name)
	}

	return list
}
