package containsduplicate

import "testing"

var cases = []struct {
	name     string
	input    []int
	expected bool
}{
	{"example_1", []int{1, 2, 3, 1}, true},
	{"example_2", []int{1, 2, 3, 4}, false},
	{"example_3", []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
}

func RunTest(t *testing.T, f func(nums []int) bool) {
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			result := f(tc.input)
			if result != tc.expected {
				t.Fatalf("expected %v, got %v", tc.expected, result)
			}
		})
	}

}
