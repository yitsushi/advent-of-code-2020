package math

// Vector is a vector interface.
type Vector interface {
	Hash() interface{}
	Neighbors() []Vector
	Values() []float64
}
