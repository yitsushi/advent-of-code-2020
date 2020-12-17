package math_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2020/pkg/math"
)

func TestVector2D_Rotate(t *testing.T) {
	tests := []struct {
		name     string
		initial  math.Vector2D
		degree   float64
		expected math.Vector2D
	}{
		{
			name:     "Rotate 90",
			initial:  math.Vector2D{X: 1, Y: 0},
			degree:   90,
			expected: math.Vector2D{X: 0, Y: 1},
		},
		{
			name:     "Rotate -90",
			initial:  math.Vector2D{X: 1, Y: 0},
			degree:   -90,
			expected: math.Vector2D{X: 0, Y: -1},
		},
		{
			name:     "Rotate 180",
			initial:  math.Vector2D{X: 1, Y: 0},
			degree:   180,
			expected: math.Vector2D{X: -1, Y: 0},
		},
		{
			name:     "Rotate -180",
			initial:  math.Vector2D{X: 1, Y: 0},
			degree:   -180,
			expected: math.Vector2D{X: -1, Y: 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &math.Vector2D{
				X: tt.initial.X,
				Y: tt.initial.Y,
			}

			v.Rotate(tt.degree)

			assert.Equal(t, tt.expected.X, v.X)
			assert.Equal(t, tt.expected.Y, v.Y)
		})
	}
}

func TestVector2D_Manhattan(t *testing.T) {
	v := math.Vector2D{15, 33}

	assert.Equal(t, float64(48), v.Manhattan())
}

func TestVector2D_Hash(t *testing.T) {
	v := math.Vector2D{15, 33}

	assert.Equal(t, "15.0000;33.0000", v.Hash())
}

func TestVector2D_Neighbors(t *testing.T) {
	v := math.Vector2D{15, 33}

	neighbors := v.Neighbors()

	assert.Len(t, neighbors, 8)
}

func TestVector2D_Values(t *testing.T) {
	v := math.Vector2D{15, 33}

	assert.Equal(t, []float64{15, 33}, v.Values())
}
