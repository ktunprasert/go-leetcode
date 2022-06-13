package longestsubstring

var cases = []struct {
	name     string
	input    string
	expected int
}{
	{"example_1", "abcabcbb", 3},
	{"example_2", "bbbbb", 1},
	{"example_3", "pwwkew", 3},
	{"gotcha_1", "aab", 2},
	{"gotcha_2", "dvdf", 3},
	{"gotcha_3", "", 0},
	{"gotcha_4", " ", 1},
	{"gotcha_4", "anviaj", 5},
}
