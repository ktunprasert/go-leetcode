package productexceptself

func productExceptSelfNaive(nums []int) []int {
	ans := make([]int, len(nums))
	for i := 0; i < len(ans); i++ {
		ans[i] = 1
	}
	var left, right int = 1, 1
	var i, j int = 0, len(nums) - 1
	for ; i < len(nums) && j >= 0; i, j = i+1, j-1 {
		ans[i] *= left
		left *= nums[i]
		ans[j] *= right
		right *= nums[j]
	}
	return ans
}
