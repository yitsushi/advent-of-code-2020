package math_test

import (
	"testing"

	"github.com/yitsushi/advent-of-code-2020/pkg/math"
)

func TestDivMod(t *testing.T) {
	type args struct {
		numerator   int64
		denominator int64
	}

	tests := []struct {
		name      string
		args      args
		quotient  int64
		remainder int64
	}{
		{
			name:      "9 divmod 3",
			args:      args{numerator: 9, denominator: 3},
			quotient:  3,
			remainder: 0,
		},
		{
			name:      "9 divmod 4",
			args:      args{numerator: 9, denominator: 4},
			quotient:  2,
			remainder: 1,
		},
		{
			name:      "9 divmod 13",
			args:      args{numerator: 9, denominator: 13},
			quotient:  0,
			remainder: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q, r := math.DivMod(tt.args.numerator, tt.args.denominator)
			if q != tt.quotient || r != tt.remainder {
				t.Errorf("DivMod() got = (%v, %v), want (%v, %v)", q, r, tt.quotient, tt.remainder)
			}
		})
	}
}
