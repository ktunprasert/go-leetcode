package productexceptself

func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))
	// first pass
	multiplier := 1
	for i := 0; i < len(ans); i++ {
		ans[i] = multiplier
		multiplier *= nums[i]
	}
	// second pass
	multiplier = 1
	for j := len(nums) - 1; j >= 0; j-- {
		ans[j] *= multiplier
		multiplier *= nums[j]
	}
	return ans
}
