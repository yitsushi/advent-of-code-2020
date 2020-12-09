package day09

import (
	"github.com/sirupsen/logrus"
	"github.com/yitsushi/advent-of-code-2020/pkg/math"
)

func (d *Solver) findPair(window, head int) (int64, error) {
	if len(d.input) <= head {
		return 0, EndOfStream{}
	}

	block := d.input[head-window : head]
	target := d.input[head]

	logrus.Debugf("Looking for %d", target)
	logrus.Debugf("In dataset of %+v", block)

	for idx, a := range block {
		for _, b := range block[idx+1:] {
			if a+b == target {
				logrus.Debugf("%d = %d + %d", target, a, b)

				return target, nil
			}
		}
	}

	return target, PairNotFound{}
}

func (d *Solver) findInvalid(window int) (int64, error) {
	head := window

	for {
		value, err := d.findPair(window, head)
		if err != nil {
			logrus.Info(value, err.Error())

			return value, err
		}

		head++
	}
}

func (d *Solver) findWeakness(invalid int64) (int64, error) {
	for i, root := range d.input[:len(d.input)-3] {
		set := []int64{root}
		sum := root

		for _, current := range d.input[i+1 : len(d.input)] {
			if sum+current > invalid {
				break
			}

			sum += current
			set = append(set, current)

			if sum == invalid {
				return math.MinimumInt64(set) + math.MaximumInt64(set), nil
			}
		}
	}

	return 0, EndOfStream{}
}
