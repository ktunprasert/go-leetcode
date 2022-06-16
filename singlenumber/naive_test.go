package singlenumber

import "testing"

type args struct {
	nums []int
}

var tests = []struct {
	name string
	args args
	want int
}{
	{"e1", args{[]int{2, 2, 1}}, 1},
	{"e2", args{[]int{4, 1, 2, 1, 2}}, 4},
	{"e3", args{[]int{1}}, 1},
}

func Test_singleNumberNaive(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumberNaive(tt.args.nums); got != tt.want {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
