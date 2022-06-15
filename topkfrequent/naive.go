package topkfrequent

func min(x, y int) int {
	if y < x {
		return y
	}
	return x
}

func topKFrequentNaive(nums []int, k int) []int {
	// Since nums slice literal gets passed in the len() call is O(1)
	if len(nums) < 2 || len(nums) == k {
		return nums
	}

	answer := make([]int, 0, k)
	freq_map := make(map[int]int)
	ans_map := make(map[int][]int)
	var max int
	for _, n := range nums {
		freq_map[n]++
		// ans_map[freq_map[n]] = []int{n}
		if ans_map[freq_map[n]] == nil {
			ans_map[freq_map[n]] = []int{n}
		} else {
			ans_map[freq_map[n]] = append(ans_map[freq_map[n]], n)
		}
		if max < freq_map[n] {
			max = freq_map[n]
		}
		// fmt.Println(ans_map)
	}
	// fmt.Println(freq_map, ans_map, max)
	for m := max; m > 0 && len(answer) < k; m-- {
		// fmt.Println(m, i, answer)
		if v, ok := ans_map[m]; ok && len(v) >= k {
			// slice_to := k - len(answer)
			// answer = append(answer, v[:min(slice_to, len(v))]...)
			answer = v[:k]
			// fmt.Println(v, answer)
		}
	}

	return answer
}
