package math

import (
	"fmt"
	"math"
)

type Vector interface {
	Hash() string
}

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

func (v *Vector2D) Hash() string {
	return fmt.Sprintf("%.4f;%.4f", v.X, v.Y)
}

// Manhattan distance.
func (v *Vector2D) Manhattan() float64 {
	return math.Abs(v.X) + math.Abs(v.Y)
}

func (v *Vector2D) Neighbors() []Vector2D {
	vectors := []Vector2D{}
	checkRange := []float64{-1, 0, 1}

	for _, x := range checkRange {
		for _, y := range checkRange {
			if x == 0 && y == 0 {
				continue
			}

			vectors = append(
				vectors,
				Vector2D{
					X: v.X + x,
					Y: v.Y + y,
				},
			)
		}
	}

	return vectors
}

// Vector3D is a 3D vector.
type Vector3D struct {
	X float64
	Y float64
	Z float64
}

func (v *Vector3D) Hash() string {
	return fmt.Sprintf("%.0f;%.0f;%.0f", v.X, v.Y, v.Z)
}

func (v *Vector3D) Neighbors() []Vector3D {
	vectors := []Vector3D{}
	checkRange := []float64{-1, 0, 1}

	for _, x := range checkRange {
		for _, y := range checkRange {
			for _, z := range checkRange {
				if x == 0 && y == 0 && z == 0 {
					continue
				}

				vectors = append(
					vectors,
					Vector3D{
						X: v.X + x,
						Y: v.Y + y,
						Z: v.Z + z,
					},
				)
			}
		}
	}

	return vectors
}

func (v *Vector3D) MinimizeFrom(vector Vector3D) {
	if vector.X < v.X {
		v.X = vector.X
	}

	if vector.Y < v.Y {
		v.Y = vector.Y
	}

	if vector.Z < v.Z {
		v.Z = vector.Z
	}
}

func (v *Vector3D) MaximizeFrom(vector Vector3D) {
	if vector.X > v.X {
		v.X = vector.X
	}

	if vector.Y > v.Y {
		v.Y = vector.Y
	}

	if vector.Z > v.Z {
		v.Z = vector.Z
	}
}
