package math

import "fmt"

// Vector3D is a 3D vector.
type Vector3D struct {
	X float64
	Y float64
	Z float64
}

// Hash for cache.
func (v Vector3D) Hash() string {
	return fmt.Sprintf("%.4f;%.4f;%.4f", v.X, v.Y, v.Z)
}

// Neighbors if it's a coordinate.
func (v Vector3D) Neighbors() []Vector {
	vectors := []Vector{}
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

// MinimizeFrom set x, y and z values to the minimum value from the two vectors.
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

// MaximizeFrom set x, y and z values to the maximum value from the two vectors.
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

// Values for the vector as plain float64 slice.
func (v Vector3D) Values() []float64 {
	return []float64{v.X, v.Y, v.Z}
}
