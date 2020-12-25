package math_test

import (
	"testing"

	"github.com/yitsushi/advent-of-code-2020/pkg/math"
)

func TestMinimumInt64(t *testing.T) {
	tests := []struct {
		name string
		list []int64
		want int64
	}{
		{name: "empty", list: []int64{}, want: 0},
		{name: "values", list: []int64{5, 4, 92, 13, 42, 52}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := math.MinimumInt64(tt.list); got != tt.want {
				t.Errorf("MinimumInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaximumInt64(t *testing.T) {
	tests := []struct {
		name string
		list []int64
		want int64
	}{
		{name: "empty", list: []int64{}, want: 0},
		{name: "values", list: []int64{5, 4, 92, 13, 42, 52}, want: 92},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := math.MaximumInt64(tt.list); got != tt.want {
				t.Errorf("MaximumInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinimumInt(t *testing.T) {
	tests := []struct {
		name string
		list []int
		want int
	}{
		{name: "empty", list: []int{}, want: 0},
		{name: "values", list: []int{5, 4, 92, 13, 42, 52}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := math.MinimumInt(tt.list); got != tt.want {
				t.Errorf("MinimumInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaximumInt(t *testing.T) {
	tests := []struct {
		name string
		list []int
		want int
	}{
		{name: "empty", list: []int{}, want: 0},
		{name: "values", list: []int{5, 4, 92, 13, 42, 52}, want: 92},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := math.MaximumInt(tt.list); got != tt.want {
				t.Errorf("MaximumInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinMaxInt(t *testing.T) {
	tests := []struct {
		name    string
		list    []int
		wantMin int
		wantMax int
	}{
		{name: "empty", list: []int{}, wantMin: 0, wantMax: 0},
		{name: "values", list: []int{5, 4, 92, 13, 42, 52}, wantMin: 4, wantMax: 92},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMin, gotMax := math.MinMaxInt(tt.list)

			if gotMin != tt.wantMin {
				t.Errorf("MinMaxInt() = Min %v, want %v", gotMin, tt.wantMin)
			}

			if gotMax != tt.wantMax {
				t.Errorf("MinMaxInt() = Max %v, want %v", gotMin, tt.wantMax)
			}
		})
	}
}
