package longestconsecutivesequence

import "testing"

func Test_longestConsecutive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"e1", args{[]int{100, 4, 200, 1, 3, 2}}, 4},
		{"e2", args{[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}}, 9},
		{"gotcha_1", args{[]int{0, -1}}, 2},
		{"gotcha_2", args{[]int{0, -1, 1}}, 3},
		{"gotcha_3", args{[]int{1, 3, 5, 2, 4}}, 5},
		{"gotcha_4", args{[]int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}}, 7},
		{"gotcha_5", args{[]int{0, 0}}, 1},
		{"self", args{[]int{-6, 0, -5, -4, -3, -2, -1}}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestConsecutive(tt.args.nums); got != tt.want {
				t.Errorf("longestConsecutive() = %v, want %v", got, tt.want)
			}
		})
	}
}
