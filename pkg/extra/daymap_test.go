package extra_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2020/pkg/extra"
	"github.com/yitsushi/advent-of-code-2020/pkg/puzzle"
)

func TestDaySelector_valid(t *testing.T) {
	day, err := extra.DaySelector(1)

	assert.NoError(t, err)
	assert.NotNil(t, day)
}

func TestDaySelector_notImplemented(t *testing.T) {
	day, err := extra.DaySelector(0)

	assert.Error(t, err)
	assert.EqualError(t, err, (puzzle.NotImplemented{}).Error())
	assert.Nil(t, day)
}

func TestDaySelector_notFound(t *testing.T) {
	day, err := extra.DaySelector(1000)

	assert.Error(t, err)
	assert.EqualError(t, err, (extra.UnkownDayError{Day: 1000}).Error())
	assert.Nil(t, day)
}

func TestDaySelector_negative(t *testing.T) {
	day, err := extra.DaySelector(-1)

	assert.Error(t, err)
	assert.EqualError(t, err, (extra.UnkownDayError{Day: -1}).Error())
	assert.Nil(t, day)
}
