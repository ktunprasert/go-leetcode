package singlenumber

func singleNumberNaive(nums []int) int {
	// optimised edge case
	if len(nums) == 1 {
		return nums[0]
	}
	var ans int
	for _, n := range nums {
		if n != 0 {
			ans ^= n
		}
	}
	return ans
}
