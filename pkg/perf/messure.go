package perf

import (
	"time"

	"github.com/sirupsen/logrus"
)

// Track starts the timer.
func Track(msg string) (string, time.Time) {
	return msg, time.Now()
}

// Duration stops the timer.
func Duration(msg string, start time.Time) {
	logrus.Infof("%v: %v\n", msg, time.Since(start))
}
