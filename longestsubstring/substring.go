package longestsubstring

func lengthOfLongestSubstring(s string) int {
	buffer := make(map[rune]int)
	var total, max int
	for _, r := range s {
		if _, ok := buffer[r]; ok {
			buffer = make(map[rune]int)
			if max < total {
				max = total
			}
			total = 0
		} else {
			buffer[r] = 1
			total++
		}
	}
	if total > max {
		return total
	}
	return max
}
