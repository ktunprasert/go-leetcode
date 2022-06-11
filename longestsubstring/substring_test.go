package longestsubstring

import (
	"fmt"
	"testing"
)

func TestLongestSubString(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected int
	}{
		{"example_1", "abcabcbb", 3},
		{"example_2", "bbbbb", 1},
		{"example_3", "pwwkew", 3},
		{"gotcha_1", "aab", 2},
		{"gotcha_2", "dvdf", 3},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			result := lengthOfLongestSubstring(tc.input)
			if tc.expected != result {
				t.Fatalf(fmt.Sprintf("Expected %d but got %d", tc.expected, result))
			}
		})
	}
}
