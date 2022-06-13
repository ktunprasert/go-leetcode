package addtwo

import (
	"fmt"
	"testing"
)

func TestAddTwoSubmission(t *testing.T) {
	for _, tc := range cases {
		t.Run(fmt.Sprintf("naive/%s", tc.name), func(t *testing.T) {
			l1 := ListNodeFromArray(tc.inputL1)
			l2 := ListNodeFromArray(tc.inputL2)
			result := addTwoNumbersEasy(l1, l2)
			if actual := fmt.Sprintf("%v", result); tc.expected != actual {
				t.Fatalf(fmt.Sprintf("Expected %s but got %s", tc.expected, actual))
			}
		})
	}
}
