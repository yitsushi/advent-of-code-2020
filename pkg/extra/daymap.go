package extra

import (
	"github.com/yitsushi/advent-of-code-2020/days/day01"
	"github.com/yitsushi/advent-of-code-2020/days/day02"
	"github.com/yitsushi/advent-of-code-2020/days/day03"
	"github.com/yitsushi/advent-of-code-2020/pkg/puzzle"
)

// DaySelector map day numbers to Day structs.
func DaySelector(day int) (puzzle.Day, error) {
	solvers := []puzzle.Day{
		nil,
		&day01.Solver{},
		&day02.Solver{},
		&day03.Solver{},
	}

	if len(solvers) <= day || day < 0 {
		return nil, puzzle.UnkownDayError{Day: day}
	}

	d := solvers[day]
	if d == nil {
		return nil, puzzle.NotImplemented{}
	}

	return d, nil
}
