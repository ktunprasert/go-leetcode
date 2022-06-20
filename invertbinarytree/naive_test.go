package invertbinarytree

import (
	"reflect"
	"testing"
)

func Test_invertTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"e1", args{BuildBinaryTree([]int{4, 2, 7, 1, 3, 6, 9}, 1)}, BuildBinaryTree([]int{4, 7, 2, 9, 6, 3, 1}, -1)},
		{"e2", args{BuildBinaryTree([]int{2, 1, 3}, 1)}, BuildBinaryTree([]int{2, 3, 1}, -1)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := invertTree(tt.args.root)
			t.Log(got, tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("invertTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
