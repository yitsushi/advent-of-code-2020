package day08

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/yitsushi/advent-of-code-2020/pkg/bullshit"
	"github.com/yitsushi/advent-of-code-2020/pkg/slice"
)

// Operation is a type alias for string.
type Operation string

// Known operations.
const (
	NOP Operation = "nop"
	ACC Operation = "acc"
	JMP Operation = "jmp"
)

// Machine is a simple machine.
type Machine struct {
	Accumulator int64
	Head        int64
	Source      []Instruction
}

// Instruction represents one instruction in the source code.
type Instruction struct {
	Operation Operation
	Value     int64
}

// NewMachine creates a new machine.
func NewMachine() Machine {
	return Machine{
		Accumulator: 0,
		Head:        0,
		Source:      []Instruction{},
	}
}

// ParseInstruction parses the string representation of an instruction.
func ParseInstruction(code string) (Instruction, error) {
	split := strings.SplitN(code, " ", 2)
	instruction := Instruction{
		Operation: NOP,
		Value:     bullshit.DropErrorInt64(strconv.ParseInt(split[1], 10, 64)),
	}

	switch split[0] {
	case "nop":
		instruction.Operation = NOP
	case "acc":
		instruction.Operation = ACC
	case "jmp":
		instruction.Operation = JMP
	default:
		return instruction, UnknownOperation{Operation: split[0], Value: instruction.Value}
	}

	return instruction, nil
}

// Load the source code into the machine.
func (m *Machine) Load(code []Instruction) {
	m.Source = code
}

// Reset the machine.
func (m *Machine) Reset() {
	m.Accumulator = 0
	m.Head = 0
}

// Step executes the next instruction.
//
// Returns true if the program gracefully terminated.
// Return false is the program did not terminate gracefully.
//
// Error contains any kind of error the machine met during the execution.
func (m *Machine) Step() (bool, error) {
	if m.Head >= int64(len(m.Source)) {
		return true, nil
	}

	current := m.Source[m.Head]

	logrus.Debugf(">>> %s", current.String())

	switch current.Operation {
	case JMP:
		m.Head += current.Value

		return false, nil
	case ACC:
		m.Accumulator += current.Value
		m.Head++

		return false, nil
	case NOP:
		m.Head++

		return false, nil
	}

	return true, ExecutionError{}
}

func (i Instruction) String() string {
	return fmt.Sprintf("%s %d", i.Operation, i.Value)
}

// Run the machine with the loaded source code.
// It runs until an error occurs or the program
// gracefully terminated.
//
// Returns true if the program gracefully terminated.
// Return false is the program did not termin\ate gracefully.
//
// Error contains any kind of error the machine met during the execution.
func (m *Machine) Run() (bool, error) {
	visitedLines := []int64{0}

	for {
		terminated, err := m.Step()
		if err != nil {
			logrus.Info(err)

			return false, err
		}

		if terminated {
			logrus.Info("exit(0)")

			return true, nil
		}

		if _, found := slice.ContainsInt64(visitedLines, m.Head); found {
			logrus.Info("Loop detected")

			return false, InfiniteLoopDetected{}
		}

		visitedLines = append(visitedLines, m.Head)
	}
}
