package day21

// Ingredient is an ingredient.
type Ingredient struct {
	Name     string
	Allergen string
}

// NewIngredient creates a new Ingredient.
func NewIngredient(name string) Ingredient {
	return Ingredient{
		Name: name,
	}
}
