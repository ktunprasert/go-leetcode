package topkfrequent

import (
	"fmt"
	"testing"
)

func TestTopKFrequent(t *testing.T) {

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			result := topKFrequent(tc.input, tc.k)
			if fmt.Sprintf("%v", result) != fmt.Sprintf("%v", tc.expected) {
				t.Fatalf("expected %v got %v", tc.expected, result)
			}
		})
	}
}
