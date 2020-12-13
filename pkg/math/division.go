package math

// DivMod is a simple div and mod together.
func DivMod(numerator, denominator int64) (int64, int64) {
	return numerator / denominator, numerator % denominator
}
