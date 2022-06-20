package invertbinarytree

import (
	"reflect"
	"testing"
)

func Test_invertTreeNaive(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := invertTreeNaive(tt.args.root)
			t.Log(got, tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("invertTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
