package topkfrequent

import (
	"fmt"
	"testing"
)

var cases = []struct {
	name     string
	input    []int
	k        int
	expected []int
}{
	{"example_1", []int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
	{"self_test", []int{5, 5, 5, 10, 9, 10}, 2, []int{5, 10}},
	{"example_2", []int{1}, 1, []int{1}},
	{"gotcha_1", []int{1, 2}, 2, []int{1, 2}},
	{"gotcha_2", []int{4, 1, -1, 2, -1, 2, 3}, 2, []int{-1, 2}},
}

func TestTopKFrequentNaive(t *testing.T) {

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			result := topKFrequentNaive(tc.input, tc.k)
			if fmt.Sprintf("%v", result) != fmt.Sprintf("%v", tc.expected) {
				t.Fatalf("expected %v got %v", tc.expected, result)
			}
		})
	}
}
