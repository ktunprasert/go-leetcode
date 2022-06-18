package minstack

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func Test_MinStack(t *testing.T) {
	type args struct {
		fns  []string
		args []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"e1",
			args{
				[]string{"MinStack", "push", "push", "push", "getMin", "pop", "top", "getMin"},
				[]string{"", "-2", "0", "-3", "", "", "", ""},
			},
			[]string{"null", "null", "null", "null", "-3", "null", "0", "-2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expects := []string{}
			stack := Constructor()
			for i, fn := range tt.args.fns {
				switch fn {
				case "MinStack":
					stack = Constructor()
					expects = append(expects, "null")
				case "push":
					val, _ := strconv.Atoi(tt.args.args[i])
					stack.Push(val)
					expects = append(expects, "null")
				case "pop":
					expects = append(expects, "null")
					stack.Pop()
				case "top":
					t.Log(stack.Top())
					expects = append(expects, fmt.Sprintf("%d", stack.Top()))
				case "getMin":
					t.Log(stack.GetMin())
					expects = append(expects, fmt.Sprintf("%d", stack.GetMin()))
				}
			}
			if !reflect.DeepEqual(expects, tt.want) {
				t.Errorf("MinStack() = %v, want %v", expects, tt.want)
			}
		})
	}
}
