package day16

// Tracker is a field tracker.
type Tracker struct {
	Index  int
	Name   string
	Values []int64
}

// NewTracker creates a new tracker.
func NewTracker(index int) *Tracker {
	return &Tracker{
		Index:  index,
		Name:   "",
		Values: []int64{},
	}
}

// AddValue saves a new known value for the field.
func (t *Tracker) AddValue(value int64) {
	t.Values = append(t.Values, value)
}

// Candidates for the represented field.
func (t *Tracker) Candidates(criteriaList []*Criteria) []*Criteria {
	candidates := []*Criteria{}

	if t.Name != "" {
		return candidates
	}

	for _, criteria := range criteriaList {
		if criteria.Taken {
			continue
		}

		if criteria.ValidateAll(t.Values) {
			candidates = append(candidates, criteria)
		}
	}

	return candidates
}

// OnMyTicket returns with the field value for my ticket.
// It's always at the end of the list.
func (t *Tracker) OnMyTicket() int64 {
	return t.Values[len(t.Values)-1]
}
