package math

import (
	"math"
)

// Vector2D is a 2D vector.
type Vector2D struct {
	X float64
	Y float64
}

const rad = math.Pi / 180

// Rotate the vector.
func (v *Vector2D) Rotate(angle float64) {
	angle *= rad

	cos, sin := math.Round(math.Cos(angle)), math.Round(math.Sin(angle))

	x := (v.X * cos) - (v.Y * sin)
	y := (v.X * sin) + (v.Y * cos)

	v.X, v.Y = x, y
}

// Manhattan distance.
func (v *Vector2D) Manhattan() float64 {
	return math.Abs(v.X) + math.Abs(v.Y)
}
