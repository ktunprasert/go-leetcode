package singlenumber

func singleNumber(nums []int) int {
	// optimised edge case
	if len(nums) == 1 {
		return nums[0]
	}
	var ans int
	for _, n := range nums {
		ans ^= n
	}
	return ans
}
