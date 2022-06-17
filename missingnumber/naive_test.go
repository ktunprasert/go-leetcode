package missingnumber

import "testing"

type args struct {
	nums []int
}

var tests = []struct {
	name string
	args args
	want int
}{
	{"e1", args{[]int{3, 0, 1}}, 2},
	{"e2", args{[]int{0, 1}}, 2},
	{"e3", args{[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}}, 8},
}

func Test_missingNumberNaive(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := missingNumberNaive(tt.args.nums); got != tt.want {
				t.Errorf("missingNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
