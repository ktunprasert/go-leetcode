package binarysearch

func search(nums []int, target int) int {
	if len(nums) == 1 && nums[0] == target {
		return 0
	}
	var cursor int = 1
	var buf int
	for cursor > 0 {
		cursor = len(nums) / 2
		switch {
		case nums[cursor] == target:
			return cursor + buf
		case nums[cursor] > target:
			nums = nums[:cursor]
		case nums[cursor] < target:
			nums = nums[cursor:]
			buf += cursor
		}
	}
	return -1
}
