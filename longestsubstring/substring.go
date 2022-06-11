package longestsubstring

// func old_lengthOfLongestSubstring(s string) int {
// 	buffer := make(map[rune]int)
// 	var total, max int
// 	for _, r := range s {
// 		if _, ok := buffer[r]; ok {
// 			buffer = make(map[rune]int)
// 			if max < total {
// 				max = total
// 			}
// 			total = 1
// 		} else {
// 			total++
// 		}
// 		buffer[r] = 1
// 	}
// 	if total > max {
// 		return total
// 	}
// 	return max
// }

func lengthOfLongestSubstring(s string) int {
	if len := len(s); len < 2 {
		return len
	}
	buffer := make(map[rune]int)
	var total, max int = 1, 0
	for i := 0; i < len(s)-1; i++ {
		l, r := rune(s[i]), rune(s[i+1])
		buffer[l] = 1
		if _, ok := buffer[r]; ok {
			buffer = make(map[rune]int)
			total = 1
			if l != r {
				buffer[l] = 1
				total = 2
			}
		} else {
			total++
		}
		if max < total {
			max = total
		}
		buffer[r] = 1
	}
	return max
}
