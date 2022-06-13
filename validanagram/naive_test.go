package validanagram

import "testing"

var cases = []struct {
	name     string
	input_s  string
	input_t  string
	expected bool
}{
	{"example_1", "anagram", "nagaram", true},
	{"example_2", "rat", "car", false},
	{"gotcha_1", "aa", "bb", false},
	{"gotcha_2", "ac", "bb", false},
}

func TestIsAnagramNaive(t *testing.T) {
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if result := isAnagram(tc.input_s, tc.input_t); tc.expected != result {
				t.Fatalf("expected %v got %v", tc.expected, result)
			}
		})
	}
}
