package day12

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/yitsushi/advent-of-code-2020/pkg/bullshit"
	"github.com/yitsushi/advent-of-code-2020/pkg/math"
)

// Solver is the main solver.
type Solver struct {
	navigator Navigator
}

const possibleCommands = "NSEWLRF"

// SetInput receives the input and parses its content.
func (d *Solver) SetInput(input io.Reader) error {
	scanner := bufio.NewScanner(input)

	d.navigator = NewNavigator(
		math.Vector2D{X: 0, Y: 0},
		math.Vector2D{X: 1, Y: 0},
	)

	for scanner.Scan() {
		value := scanner.Text()
		if len(value) <= 1 {
			return UnknownInstruction{value}
		}

		instruction := Instruction{
			Command: value[0],
			Value:   float64(bullshit.DropErrorInt64(strconv.ParseInt(value[1:], 10, 64))),
		}
		if !strings.ContainsRune(possibleCommands, rune(instruction.Command)) {
			return UnknownInstruction{value}
		}

		d.navigator.AddInstruction(instruction)
	}

	return scanner.Err()
}
