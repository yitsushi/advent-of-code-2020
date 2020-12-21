package day20_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2020/days/day20"
)

func TestTile_Rotate_1(t *testing.T) {
	tile := day20.NewTile(1)
	tile.Data = []string{
		"qwe",
		"asd",
		"zxc",
	}

	expected := []string{
		"edc",
		"wsx",
		"qaz",
	}

	tile.Rotate(1)

	assert.Equal(t, expected, tile.Data)
}

func TestTile_Rotate_2(t *testing.T) {
	tile := day20.NewTile(2)
	tile.Data = []string{
		"qwe",
		"asd",
		"zxc",
	}

	expected := []string{
		"cxz",
		"dsa",
		"ewq",
	}

	tile.Rotate(2)

	assert.Equal(t, expected, tile.Data)
}

func TestTile_Rotate_3(t *testing.T) {
	tile := day20.NewTile(3)
	tile.Data = []string{
		"qwe",
		"asd",
		"zxc",
	}

	expected := []string{
		"zaq",
		"xsw",
		"cde",
	}

	tile.Rotate(3)

	assert.Equal(t, expected, tile.Data)
}

func TestTile_Rotate_12(t *testing.T) {
	tile := day20.NewTile(12)
	tile.Data = []string{
		"qwe",
		"asd",
		"zxc",
	}

	expected := []string{
		"qwe",
		"asd",
		"zxc",
	}

	tile.Rotate(12)

	assert.Equal(t, expected, tile.Data)
}
