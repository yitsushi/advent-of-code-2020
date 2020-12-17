package math_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2020/pkg/math"
)

func TestVector3D_Hash(t *testing.T) {
	v := math.Vector3D{15, 33, 11}

	assert.Equal(t, "15.0000;33.0000;11.0000", v.Hash())
}

func TestVector3D_Neighbors(t *testing.T) {
	v := math.Vector3D{15, 33, 11}

	neighbors := v.Neighbors()

	assert.Len(t, neighbors, 26)
}

func TestVector3D_MinimizeFrom(t *testing.T) {
	v := math.Vector3D{15, 33, 11}
	v2 := math.Vector3D{33, 12, 13}
	expected := math.Vector3D{15, 12, 11}

	v.MinimizeFrom(v2)

	assert.Equal(t, expected, v)
}

func TestVector3D_MaximizeFrom(t *testing.T) {
	v := math.Vector3D{15, 33, 11}
	v2 := math.Vector3D{33, 12, 13}
	expected := math.Vector3D{33, 33, 13}

	v.MaximizeFrom(v2)

	assert.Equal(t, expected, v)
}

func TestVector3D_MinimizeFrom_full(t *testing.T) {
	v := math.Vector3D{33, 44, 13}
	v2 := math.Vector3D{15, 33, 11}
	expected := math.Vector3D{15, 33, 11}

	v.MinimizeFrom(v2)

	assert.Equal(t, expected, v)
}

func TestVector3D_MaximizeFrom_full(t *testing.T) {
	v := math.Vector3D{10, 12, 8}
	v2 := math.Vector3D{15, 33, 11}
	expected := math.Vector3D{15, 33, 11}

	v.MaximizeFrom(v2)

	assert.Equal(t, expected, v)
}
