package groupanagrams

import (
	"fmt"
	"testing"
)

var cases = []struct {
	name     string
	input    []string
	expected [][]string
}{
	{"example_1", []string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{[]string{"bat"}, []string{"nat", "tan"}, []string{"ate", "eat", "tea"}}},
	{"example_2", []string{""}, [][]string{[]string{""}}},
}

func TestGroupAnagramsNaive(t *testing.T) {
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			result := groupAnagramsNaive(tc.input)
			if fmt.Sprintf("%v", result) != fmt.Sprintf("%v", tc.expected) {
				t.Fatalf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}
