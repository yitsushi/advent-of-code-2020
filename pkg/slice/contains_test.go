package slice_test

import (
	"testing"

	"github.com/yitsushi/advent-of-code-2020/pkg/slice"
)

func TestContainsInt64(t *testing.T) {
	type args struct {
		list  []int64
		value int64
	}

	tests := []struct {
		name string
		args args
		want int
		ok   bool
	}{
		{
			name: "in list",
			args: args{list: []int64{3, 6, 9, 2, 3, 8}, value: 9},
			want: 2,
			ok:   true,
		},
		{
			name: "not in list",
			args: args{list: []int64{3, 6, 9, 2, 3, 8}, value: 13},
			want: -1,
			ok:   false,
		},
		{
			name: "in list repeated",
			args: args{list: []int64{3, 6, 9, 2, 3, 8}, value: 3},
			want: 0,
			ok:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := slice.ContainsInt64(tt.args.list, tt.args.value)

			if got != tt.want || ok != tt.ok {
				t.Errorf("ContainsInt64() = (%v, %v), want (%v, %v)", got, ok, tt.want, tt.ok)
			}
		})
	}
}
