package longestsubstring

import (
	"fmt"
	"testing"
)

func TestLongestSubStringNaive(t *testing.T) {
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			result := lengthOfLongestSubstringNaive(tc.input)
			if tc.expected != result {
				t.Fatalf(fmt.Sprintf("Expected %d but got %d", tc.expected, result))
			}
		})
	}
}
