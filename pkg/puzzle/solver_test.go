package puzzle_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2020/pkg/puzzle"
)

func TestSolver_Solve_valid(t *testing.T) {
	solver := puzzle.NewSolver(&testDay{}, 1)
	solution, err := solver.Solve("something\nelse", "")

	assert.NoError(t, err)
	assert.Equal(t, "yey", solution)
}

func TestSolver_Solve_validNoSolution(t *testing.T) {
	solver := puzzle.NewSolver(&testDay{}, 1)
	solution, err := solver.Solve("value", "")

	assert.Error(t, err)
	assert.Equal(t, "2020404 - solution not found", err.Error())
	assert.Equal(t, "", solution)
}

func TestSolver_Solve_notImplementedPart2(t *testing.T) {
	solver := puzzle.NewSolver(&testDay{}, 2)
	solution, err := solver.Solve("value", "")

	assert.Error(t, err)
	assert.Equal(t, "2020501 - not implemented", err.Error())
	assert.Equal(t, "", solution)
}

func TestSolver_Solve_part1AndnotImplementedPart2(t *testing.T) {
	solver := puzzle.NewSolver(&testDay{}, 0)
	solution, err := solver.Solve("something\nelse", "")

	assert.Error(t, err)
	assert.Equal(t, "2020501 - not implemented", err.Error())
	assert.Equal(t, "Part1: yey\n\nPart2: ", solution)
}

func TestSolver_Solve_part1FailNoPart2Called(t *testing.T) {
	solver := puzzle.NewSolver(&testDay{}, 0)
	solution, err := solver.Solve("value", "")

	assert.Error(t, err)
	assert.Equal(t, "2020404 - solution not found", err.Error())
	assert.Equal(t, "", solution)
}

func TestSolver_Solve_invalidInput(t *testing.T) {
	solver := puzzle.NewSolver(&testDay{}, 1)
	solution, err := solver.Solve("nope", "")

	assert.Error(t, err)
	assert.Equal(t, "", solution)
}

func TestSolver_Solve_noFile(t *testing.T) {
	solver := puzzle.NewSolver(&testDay{}, 1)
	solution, err := solver.Solve("", "nothing")

	assert.Error(t, err)
	assert.Equal(t, "", solution)
}

func TestSolver_Solve_invalidPart(t *testing.T) {
	solver := puzzle.NewSolver(&testDay{}, 8)
	solution, err := solver.Solve("value", "")

	assert.Error(t, err)
	assert.Equal(t, "unknown part: *puzzle_test.testDay :: 8", err.Error())
	assert.Equal(t, "", solution)
}

func TestSolver_Solve_noInput(t *testing.T) {
	solver := puzzle.NewSolver(&testDay{}, 8)
	solution, err := solver.Solve("", "")

	assert.Error(t, err)
	assert.Equal(t, "no input, no puzzle", err.Error())
	assert.Equal(t, "", solution)
}
