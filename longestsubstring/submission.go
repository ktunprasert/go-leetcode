package longestsubstring

func lengthOfLongestSubstring(s string) int {
	seen := make(map[rune]int)
	var max, left int
	for i, char := range s {
		switch seen_idx, ok := seen[char]; {
		case ok && !(seen_idx < left):
			left = seen_idx + 1
		default:
			if len := i - left + 1; max < len {
				max = len
			}
		}
		seen[char] = i
	}

	return max
}
