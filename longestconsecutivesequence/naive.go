package longestconsecutivesequence

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestConsecutive(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	evaluated := make(map[int]*struct{})
	for _, n := range nums {
		evaluated[n] = &struct{}{}
	}

	var maximumSequence int
	for key, _ := range evaluated {
		if evaluated[key-1] != nil {
			continue
		}
		var count int
		for ; evaluated[key] != nil; key++ {
			count++
		}
		maximumSequence = max(maximumSequence, count)
	}

	return maximumSequence
}
