package missingnumber

func missingNumberNaive(nums []int) int {
	var missing int
	for i := 0; i < len(nums); i++ {
		missing ^= nums[i] ^ (i + 1)
	}
	return missing
}
