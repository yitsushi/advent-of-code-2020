package day11_test

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2020/days/day11"
	"github.com/yitsushi/advent-of-code-2020/test"
)

func loadFerry(fixture string) (day11.Ferry, error) {
	example, err := test.LoadFixture(fixture)
	if err != nil {
		return day11.Ferry{}, err
	}

	scanner := bufio.NewScanner(ioutil.NopCloser(bytes.NewReader(example)))
	ferry := day11.NewFerry()

	for scanner.Scan() {
		row, err := day11.ParseRow(scanner.Text())
		if err != nil {
			return ferry, err
		}

		ferry.AddRow(row)
	}

	return ferry, scanner.Err()
}

// Short

func TestFerry_Adjacents_short1(t *testing.T) {
	ferry, err := loadFerry("adjacents1")
	if !assert.NoError(t, err) {
		return
	}

	adjacents := ferry.Adjacents(3, 4, false)

	expected, err := day11.ParseRow("...#...#")
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, expected, day11.Row(adjacents))
}

func TestFerry_Adjacents_short2(t *testing.T) {
	ferry, err := loadFerry("adjacents2")
	if !assert.NoError(t, err) {
		return
	}

	adjacents := ferry.Adjacents(1, 1, false)

	expected, err := day11.ParseRow("........")
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, expected, day11.Row(adjacents))
}

func TestFerry_Adjacents_short3(t *testing.T) {
	ferry, err := loadFerry("adjacents3")
	if !assert.NoError(t, err) {
		return
	}

	adjacents := ferry.Adjacents(3, 3, false)

	expected, err := day11.ParseRow("........")
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, expected, day11.Row(adjacents))
}

// Long

func TestFerry_Adjacents_long1(t *testing.T) {
	ferry, err := loadFerry("adjacents1")
	if !assert.NoError(t, err) {
		return
	}

	adjacents := ferry.Adjacents(3, 4, true)

	expected, err := day11.ParseRow("########")
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, expected, day11.Row(adjacents))
}

func TestFerry_Adjacents_long2(t *testing.T) {
	ferry, err := loadFerry("adjacents2")
	if !assert.NoError(t, err) {
		return
	}

	adjacents := ferry.Adjacents(1, 1, true)

	expected, err := day11.ParseRow("....L...")
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, expected, day11.Row(adjacents))
}

func TestFerry_Adjacents_long3(t *testing.T) {
	ferry, err := loadFerry("adjacents3")
	if !assert.NoError(t, err) {
		return
	}

	adjacents := ferry.Adjacents(3, 3, true)

	expected, err := day11.ParseRow("........")
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, expected, day11.Row(adjacents))
}
