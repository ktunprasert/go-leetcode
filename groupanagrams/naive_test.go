package groupanagrams

import (
	"fmt"
	"sort"
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

func sortByLength(res [][]string) [][]string {
	// pass-by-value so mutation wouldn't matter
	sort.Slice(res, func(i, j int) bool { return len(res[i]) < len(res[j]) })
	return res
}

func sortByFirstChar(res []string) []string {
	sort.Slice(res, func(i, j int) bool { return res[i] < res[j] })
	return res
}

func TestGroupAnagramsNaive(t *testing.T) {
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			out := groupAnagramsNaive(tc.input)
			sorted := sortByLength(out)
			final := [][]string{}
			for _, str_arr := range sorted {
				final = append(final, sortByFirstChar(str_arr))
			}
			if fmt.Sprintf("%v", final) != fmt.Sprintf("%v", tc.expected) {
				t.Fatalf("expected %v, got %v", tc.expected, final)
			}
		})
	}
}
