package day17_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2020/days/day17"
	"github.com/yitsushi/advent-of-code-2020/pkg/math"
)

type spaceCheck struct {
	coordinate math.Vector3D
	isNil      bool
	active     bool
}

func TestSpace(t *testing.T) {
	space := day17.NewSpace()

	space.Inspect(math.Vector3D{X: 0, Y: 0, Z: 0})
	space.Inspect(math.Vector3D{X: 1, Y: 0, Z: 0})
	space.Inspect(math.Vector3D{X: 1, Y: -1, Z: -1}).Activate()
	space.Inspect(math.Vector3D{X: 1, Y: 0, Z: 0}).Activate()

	space.Finalize()

	testCases := []spaceCheck{
		{
			coordinate: math.Vector3D{X: 1, Y: -1, Z: -1},
			isNil:      false,
			active:     true,
		},
		{
			coordinate: math.Vector3D{X: 0, Y: 0, Z: 0},
			isNil:      false,
			active:     false,
		},
		{
			coordinate: math.Vector3D{X: -1, Y: 0, Z: -1},
			isNil:      true,
			active:     false,
		},
	}

	for _, testCase := range testCases {
		node := space.Lookup(testCase.coordinate)

		if testCase.isNil {
			assert.Nil(t, node)
		} else {
			assert.NotNil(t, node)
			assert.Equal(t, testCase.active, node.Active)
		}

		node = space.Inspect(testCase.coordinate)

		assert.NotNil(t, node)
		assert.Equal(t, testCase.active, node.Active)
		assert.Equal(t, testCase.coordinate.Hash(), node.Coordinate.Hash())
	}

	node := space.Inspect(math.Vector3D{X: 0, Y: 0, Z: 0})

	assert.Len(t, space.Neighbors(node), 26)
	assert.Len(t, space.SelectNeighbors(node, true), 2)
	assert.Len(t, space.SelectNeighbors(node, false), 24)

	node.Dectivate()
	node.Finalize()

	assert.False(t, node.Active)

	node.ToggleActive()
	node.Finalize()

	assert.True(t, node.Active)

	notThere := math.Vector3D{X: 10, Y: 10, Z: 10}

	assert.Len(t, space.Neighbors(space.Lookup(notThere)), 0)
	assert.Len(t, space.SelectNeighbors(space.Lookup(notThere), true), 0)
}
