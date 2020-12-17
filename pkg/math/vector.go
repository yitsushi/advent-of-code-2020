package math

// Vector is a vector interface.
type Vector interface {
	Hash() string
	Neighbors() []Vector
	Values() []float64
}
