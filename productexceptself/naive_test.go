package productexceptself

import (
	"reflect"
	"testing"
)

type args struct {
	nums []int
}

var tests = []struct {
	name string
	args args
	want []int
}{
	{"example_1", args{[]int{1, 2, 3, 4}}, []int{24, 12, 8, 6}},
	{"example_2", args{[]int{-1, 1, 0, -3, 3}}, []int{0, 0, 9, 0, 0}},
}

func Test_productExceptSelfNaive(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := productExceptSelfNaive(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("productExceptSelf() = %v, want %v", got, tt.want)
			}
		})
	}
}
