package day16

// Criteria is a rule defined in the input file.
type Criteria struct {
	Name   string
	Ranges []Range
	Taken  bool
}

// Validate returns true if the value matches the criteria.
func (c *Criteria) Validate(value int64) bool {
	for _, r := range c.Ranges {
		if value >= r.From && value <= r.To {
			return true
		}
	}

	return false
}

// ValidateAll returns true if all values are valid for this criteria.
func (c *Criteria) ValidateAll(values []int64) bool {
	for _, value := range values {
		if !c.Validate(value) {
			return false
		}
	}

	return true
}

// Range for Criteria.
type Range struct {
	From int64
	To   int64
}

// Ticket is a single ticket with unknown fields.
type Ticket struct {
	Fields []int64
}

// Validate a ticket with the provided criteria list.
func (t *Ticket) Validate(criteriaList []*Criteria) []int64 {
	invalidValues := []int64{}

	for _, value := range t.Fields {
		valid := false

		for _, criteria := range criteriaList {
			if criteria.Validate(value) {
				valid = true

				break
			}
		}

		if !valid {
			invalidValues = append(invalidValues, value)
		}
	}

	return invalidValues
}
