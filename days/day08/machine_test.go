package day08_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/advent-of-code-2020/days/day08"
)

func TestMachine_Run(t *testing.T) {
	machine := day08.NewMachine()

	machine.Load([]day08.Instruction{
		{Operation: day08.ACC, Value: 10},
		{Operation: day08.NOP, Value: 1},
	})

	terminated, err := machine.Run()

	assert.True(t, terminated)
	assert.NoError(t, err)
}

func TestMachine_Run_loop(t *testing.T) {
	machine := day08.NewMachine()

	machine.Load([]day08.Instruction{
		{Operation: day08.ACC, Value: 10},
		{Operation: day08.NOP, Value: 1},
		{Operation: day08.JMP, Value: -1},
	})

	terminated, err := machine.Run()

	assert.False(t, terminated)
	assert.Error(t, err)
	assert.Equal(t, "infinite loop detected", err.Error())
}

func TestMachine_Run_unknownOperation(t *testing.T) {
	machine := day08.NewMachine()

	machine.Load([]day08.Instruction{
		{Operation: "something", Value: 10},
	})

	terminated, err := machine.Run()

	assert.False(t, terminated)
	assert.Error(t, err)
	assert.Equal(t, "execution error", err.Error())
}
