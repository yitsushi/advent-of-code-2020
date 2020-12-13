package math_test

import (
	"testing"

	"github.com/yitsushi/advent-of-code-2020/pkg/math"
)

func TestLCMSlice(t *testing.T) {
	tests := []struct {
		name string
		args []int64
		want int64
	}{
		{
			name: "example 1",
			args: []int64{1, 2, 3, 4, 28},
			want: 84,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := math.LCMSlice(tt.args); got != tt.want {
				t.Errorf("LCMSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
