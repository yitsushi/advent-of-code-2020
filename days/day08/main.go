package day08

import (
	"fmt"

	"github.com/yitsushi/advent-of-code-2020/pkg/perf"
	"github.com/yitsushi/advent-of-code-2020/pkg/puzzle"
)

// Part1 for this day.
func (d *Solver) Part1() (string, error) {
	defer perf.Duration(perf.Track("Part1"))

	terminated, err := d.machine.Run()
	if terminated {
		return "", puzzle.NoSolution{}
	}

	switch err.(type) {
	case InfiniteLoopDetected:
		return fmt.Sprintf("%d", d.machine.Accumulator), nil
	default:
		return "", err
	}
}

// Part2 for this day.
func (d *Solver) Part2() (string, error) {
	defer perf.Duration(perf.Track("Part2"))

	replaceMap := map[Operation]Operation{
		JMP: NOP,
		NOP: JMP,
	}

	for idx, instruction := range d.machine.Source {
		if v, ok := replaceMap[instruction.Operation]; ok {
			instruction.Operation = v
		} else {
			continue
		}

		machine := NewMachine()
		machine.Source = make([]Instruction, len(d.machine.Source))

		copy(machine.Source, d.machine.Source)

		machine.Source[idx] = instruction

		terminated, err := machine.Run()
		if terminated {
			return fmt.Sprintf("%d", machine.Accumulator), nil
		}

		switch err.(type) {
		case InfiniteLoopDetected:
			continue
		default:
			return "", err
		}
	}

	return "", puzzle.NotImplemented{}
}
