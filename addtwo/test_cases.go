package addtwo

var cases = []struct {
	name     string
	inputL1  []int
	inputL2  []int
	expected string
}{
	{"same_size_arr", []int{2, 4, 3}, []int{5, 6, 4}, "[7 0 8]"},
	{"zero", []int{0}, []int{0}, "[0]"},
	{"uneven", []int{0, 1}, []int{3}, "[3 1]"},
	{"example_3", []int{9, 9, 9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}, "[8 9 9 9 0 0 0 0 0 1]"},
}
