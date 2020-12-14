package extra

import (
	"github.com/yitsushi/advent-of-code-2020/days/day01"
	"github.com/yitsushi/advent-of-code-2020/days/day02"
	"github.com/yitsushi/advent-of-code-2020/days/day03"
	"github.com/yitsushi/advent-of-code-2020/days/day04"
	"github.com/yitsushi/advent-of-code-2020/days/day05"
	"github.com/yitsushi/advent-of-code-2020/days/day06"
	"github.com/yitsushi/advent-of-code-2020/days/day07"
	"github.com/yitsushi/advent-of-code-2020/days/day08"
	"github.com/yitsushi/advent-of-code-2020/days/day09"
	"github.com/yitsushi/advent-of-code-2020/days/day10"
	"github.com/yitsushi/advent-of-code-2020/days/day11"
	"github.com/yitsushi/advent-of-code-2020/days/day12"
	"github.com/yitsushi/advent-of-code-2020/days/day13"
	"github.com/yitsushi/advent-of-code-2020/days/day14"
	"github.com/yitsushi/advent-of-code-2020/days/day15"
	"github.com/yitsushi/advent-of-code-2020/days/day16"
	"github.com/yitsushi/advent-of-code-2020/days/day17"
	"github.com/yitsushi/advent-of-code-2020/days/day18"
	"github.com/yitsushi/advent-of-code-2020/days/day19"
	"github.com/yitsushi/advent-of-code-2020/days/day20"
	"github.com/yitsushi/advent-of-code-2020/days/day21"
	"github.com/yitsushi/advent-of-code-2020/days/day22"
	"github.com/yitsushi/advent-of-code-2020/days/day23"
	"github.com/yitsushi/advent-of-code-2020/days/day24"
	"github.com/yitsushi/advent-of-code-2020/pkg/puzzle"
)

// DaySelector map day numbers to Day structs.
func DaySelector(day int) (puzzle.Day, error) {
	solvers := []puzzle.Day{
		nil,
		&day01.Solver{},
		&day02.Solver{},
		&day03.Solver{},
		&day04.Solver{},
		&day05.Solver{},
		&day06.Solver{},
		&day07.Solver{},
		&day08.Solver{},
		&day09.Solver{},
		&day10.Solver{},
		&day11.Solver{},
		&day12.Solver{},
		&day13.Solver{},
		&day14.Solver{},
		&day15.Solver{},
		&day16.Solver{},
		&day17.Solver{},
		&day18.Solver{},
		&day19.Solver{},
		&day20.Solver{},
		&day21.Solver{},
		&day22.Solver{},
		&day23.Solver{},
		&day24.Solver{},
	}

	if len(solvers) <= day || day < 0 {
		return nil, UnkownDayError{Day: day}
	}

	d := solvers[day]
	if d == nil {
		return nil, puzzle.NotImplemented{}
	}

	return d, nil
}
