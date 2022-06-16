package singlenumber

func singleNumber(nums []int) int {
	var ans int = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] != 0 {
			ans ^= nums[i]
		}
	}
	return ans
}
