package plusone

func plusOne(digits []int) []int {
	var carry int

	sum := digits[len(digits)-1] + 1
	digits[len(digits)-1] = sum % 10
	carry = sum / 10

	for i := len(digits) - 2; i >= 0 && carry > 0; i-- {
		sum := digits[i] + carry
		carry = sum / 10
		digits[i] = sum % 10
	}

	if carry != 0 {
		return append([]int{carry}, digits...)
	}

	return digits
}
