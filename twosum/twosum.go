package twosum

func twoSum(nums []int, target int) []int {
	buffer := make(map[int]int, len(nums))
	for i, n := range nums {
		if buffer_i, ok := buffer[target-n]; ok {
			return []int{buffer_i, i}
		}
		buffer[n] = i
	}
	return []int{-1, -1}
}
