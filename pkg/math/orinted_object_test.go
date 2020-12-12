package math_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2020/pkg/math"
)

func TestOrintedObject2D_Move(t *testing.T) {
	object := math.OrintedObject2D{
		Position: math.Vector2D{15, 33},
		Facing:   math.Vector2D{3, 1},
	}

	object.Move(10)

	assert.Equal(t, float64(45), object.Position.X)
	assert.Equal(t, float64(43), object.Position.Y)
}

func TestOrintedObject2D_Rotate(t *testing.T) {
	object := math.OrintedObject2D{
		Position: math.Vector2D{15, 33},
		Facing:   math.Vector2D{3, 1},
	}

	object.Rotate(90)

	assert.Equal(t, float64(-1), object.Facing.X)
	assert.Equal(t, float64(3), object.Facing.Y)
}
