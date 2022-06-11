package twosum

import (
	"testing"
)

func TestTwoSumFirstCase(t *testing.T) {
	ans := twoSum([]int{2, 7, 11, 15}, 9)
	expected := [2]int{0, 1}
	if ans[0] != expected[0] && ans[1] != expected[1] {
		t.Fatalf("incorrect answer")
	}
}
