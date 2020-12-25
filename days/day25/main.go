package day25

import (
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/yitsushi/advent-of-code-2020/pkg/perf"
)

const (
	subjectNumber = 7
	modulo        = 20201227
)

// Part1 for this day.
func (d *Solver) Part1() (string, error) {
	defer perf.Duration(perf.Track("Part1"))

	loopNumber, match := calculateLoopNumber(d.input[0], d.input[1])
	target := d.input[0]

	if match == d.input[0] {
		target = d.input[1]
	}

	var encryptionKey int64 = 1

	for ; loopNumber > 0; loopNumber-- {
		encryptionKey *= target
		encryptionKey %= modulo
	}

	logrus.Info(loopNumber, encryptionKey)

	return strconv.FormatInt(encryptionKey, 10), nil
}

// Part2 for this day.
func (d *Solver) Part2() (string, error) {
	return "Happy Yule!", nil
}

func calculateLoopNumber(pubKey1, pubKey2 int64) (int64, int64) {
	var (
		loopNumber int64 = 0
		value      int64 = 1
	)

	for ; value != pubKey1 && value != pubKey2; loopNumber++ {
		value *= subjectNumber
		value %= modulo
	}

	match := pubKey1

	if value == pubKey2 {
		match = pubKey2
	}

	return loopNumber, match
}
