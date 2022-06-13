package containsduplicate

func containsDuplicateHashMap(nums []int) bool {
	buffer := make(map[int]struct{})
	for _, n := range nums {
		if _, ok := buffer[n]; ok {
			return true
		}
		buffer[n] = struct{}{}
	}
	return false
}
