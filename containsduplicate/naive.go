package containsduplicate

func containsDuplicateHashMap(nums []int) bool {
	buffer := make(map[int]bool)
	for _, n := range nums {
		switch _, ok := buffer[n]; {
		case ok:
			return true
		default:
			buffer[n] = true
		}
	}
	return false
}
