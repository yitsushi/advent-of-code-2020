package day14

import (
	"fmt"

	"github.com/yitsushi/advent-of-code-2020/pkg/perf"
)

func (d *Solver) sum() uint64 {
	sum := uint64(0)

	for _, v := range d.memory {
		sum += v
	}

	return sum
}

// Part1 for this day.
func (d *Solver) Part1() (string, error) {
	defer perf.Duration(perf.Track("Part1"))

	for _, op := range d.operations {
		if op.IsMask() {
			d.currentMask = op.Mask

			continue
		}

		d.memory[op.Address] = (op.Value & d.currentMask.And) | d.currentMask.Or
	}

	return fmt.Sprintf("%d", d.sum()), nil
}

// Part2 for this day.
func (d *Solver) Part2() (string, error) {
	defer perf.Duration(perf.Track("Part2"))

	for _, op := range d.operations {
		if op.IsMask() {
			d.currentMaskv2 = op.Maskv2

			continue
		}

		address := op.Address | d.currentMaskv2.Or

		for i := 0; i < 1<<len(d.currentMaskv2.Floating); i++ {
			for idx, m := range d.currentMaskv2.Floating {
				if i&(1<<idx) > 0 {
					address |= m.Or
				} else {
					address &= m.And
				}
			}

			d.memory[address] = op.Value
		}
	}

	return fmt.Sprintf("%d", d.sum()), nil
}
