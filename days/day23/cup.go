package day23

// Cup is one cup.
type Cup struct {
	Value int
	Next  *Cup
}

// Nth cup clockwise.
func (cup *Cup) Nth(n int) *Cup {
	target := cup

	for ; n > 0; n-- {
		target = target.Next
	}

	return target
}

// Find cup with label.
func (cup *Cup) Find(n int) *Cup {
	target := cup.Next

	for {
		if target.Value == n {
			return target
		}

		if target == cup {
			return nil
		}

		target = target.Next
	}
}
