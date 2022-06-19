package binarysearch

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"e1", args{[]int{-1, 0, 3, 5, 9, 12}, 9}, 4},
		{"e2", args{[]int{-1, 0, 3, 5, 9, 12}, 2}, -1},
		{"self1", args{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 3}, 3},
		{"self2", args{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 9}, 9},
		{"gotcha1", args{[]int{5}, -5}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
