package topkfrequent

func topKFrequent(nums []int, k int) []int {
	// Since nums slice literal gets passed in the len() call is O(1)
	if len(nums) < 2 || len(nums) == k {
		return nums
	}

	freq_map := make(map[int]int)
	ans_map := make(map[int][]int)
	for _, n := range nums {
		freq_map[n]++
		if ans_map[freq_map[n]] == nil {
			ans_map[freq_map[n]] = []int{n}
		} else {
			ans_map[freq_map[n]] = append(ans_map[freq_map[n]], n)
		}
		// fmt.Println(ans_map)
	}

	for i := len(ans_map); ; i-- {
		if len(ans_map[i]) >= k {
			// fmt.Println("Hi", ans_map[i])
			return ans_map[i][:k]
		}
	}
}
