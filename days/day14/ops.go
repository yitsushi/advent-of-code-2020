package day14

// Mask is a simple mask that has an And and an Or mask.
// Both of them will be applied in Part 1.
type Mask struct {
	And uint64
	Or  uint64
}

// Maskv2 is the Version 2 Masking Chip.
// Or applied always.
// All floating masks are applies in all combinations.
type Maskv2 struct {
	Or       uint64
	Floating []Mask
}

// Operation represents one operation.
type Operation struct {
	Address uint64
	Value   uint64
	Mask    *Mask
	Maskv2  *Maskv2
}

// IsMask is a simple helper function, so I don't have
// to write ugly `if operation.Mask != nil`.
func (o *Operation) IsMask() bool {
	return o.Mask != nil
}
