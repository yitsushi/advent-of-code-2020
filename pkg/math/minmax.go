package math

// MinimumInt64 value of a slice.
func MinimumInt64(list []int64) int64 {
	if len(list) < 1 {
		return 0
	}

	min := list[0]

	for _, value := range list[1:] {
		if value < min {
			min = value
		}
	}

	return min
}

// MaximumInt64 value of a slice.
func MaximumInt64(list []int64) int64 {
	if len(list) < 1 {
		return 0
	}

	max := list[0]

	for _, value := range list[1:] {
		if value > max {
			max = value
		}
	}

	return max
}

// MinimumInt value of a slice.
func MinimumInt(list []int) int {
	if len(list) < 1 {
		return 0
	}

	min := list[0]

	for _, value := range list[1:] {
		if value < min {
			min = value
		}
	}

	return min
}

// MaximumInt value of a slice.
func MaximumInt(list []int) int {
	if len(list) < 1 {
		return 0
	}

	max := list[0]

	for _, value := range list[1:] {
		if value > max {
			max = value
		}
	}

	return max
}

// MinMaxInt value of a slice.
func MinMaxInt(list []int) (int, int) {
	if len(list) < 1 {
		return 0, 0
	}

	max, min := list[0], list[0]

	for _, value := range list[1:] {
		if value > max {
			max = value
		}

		if value < min {
			min = value
		}
	}

	return min, max
}
