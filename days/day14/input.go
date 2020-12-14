package day14

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strings"
)

const (
	bitmaskFor36bits = 0xfffffffff
)

// Solver is the main solver.
type Solver struct {
	currentMask   *Mask
	currentMaskv2 *Maskv2
	memory        map[uint64]uint64
	operations    []Operation
}

// SetInput receives the input and parses its content.
func (d *Solver) SetInput(input io.Reader) error {
	scanner := bufio.NewScanner(input)
	d.memory = map[uint64]uint64{}
	d.operations = []Operation{}

	for scanner.Scan() {
		text := scanner.Text()

		op := d.parseMask(text)
		if op == nil {
			op = d.parseAssignment(text)
		}

		if op != nil {
			d.operations = append(d.operations, *op)
		}
	}

	return scanner.Err()
}

func (d *Solver) parseAssignment(text string) *Operation {
	if !strings.HasPrefix(text, "mem") {
		return nil
	}

	var op Operation

	fmt.Sscanf(text, "mem[%d] = %d", &op.Address, &op.Value)

	return &op
}

func (d *Solver) parseMask(text string) *Operation {
	if !strings.HasPrefix(text, "mask = ") {
		return nil
	}

	var (
		stream string
		op     Operation
	)

	fmt.Sscanf(text, "mask = %s", &stream)

	op.Mask = &Mask{And: math.MaxUint64 & bitmaskFor36bits, Or: 0}
	op.Maskv2 = &Maskv2{Or: 0, Floating: []Mask{}}

	l := len(stream) - 1

	for idx := l; idx >= 0; idx-- {
		bit := stream[idx]
		v := uint64(1 << (l - idx))

		switch bit {
		case '1':
			op.Mask.Or ^= v
			op.Maskv2.Or ^= v
		case '0':
			op.Mask.And &= ^v
		case 'X':
			op.Maskv2.Floating = append(op.Maskv2.Floating, Mask{
				And: math.MaxUint64 & bitmaskFor36bits & (^v),
				Or:  v,
			})
		}
	}

	return &op
}
