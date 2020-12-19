package slice

// EqualInt64 test if two slices are equal or not.
// They are equal if the yhave the same set of element
// in the same order.
func EqualInt64(a, b []int64) bool {
	if len(a) != len(b) {
		return false
	}

	for idx := 0; idx < len(a); idx++ {
		if a[idx] != b[idx] {
			return false
		}
	}

	return true
}
