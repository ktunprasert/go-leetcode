package groupanagrams

import (
	"fmt"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	for _, tc := range cases {
		out := groupAnagramsNaive(tc.input)
		sorted := sortByLength(out)
		final := [][]string{}
		for _, str_arr := range sorted {
			final = append(final, sortByFirstChar(str_arr))
		}
		if fmt.Sprintf("%v", final) != fmt.Sprintf("%v", tc.expected) {
			t.Fatalf("expected %v, got %v", tc.expected, final)
		}

	}
}
