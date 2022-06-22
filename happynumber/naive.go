package happynumber

func isHappy(n int) bool {
	evaluated := make(map[int]struct{})
	var sum int
	for n != 1 {
		sum = 0
		for i := n; i > 0; i /= 10 {
			sum += (i % 10) * (i % 10)
		}

		if _, ok := evaluated[sum]; ok {
			break
		}

		evaluated[sum] = struct{}{}
		n = sum
	}

	return n == 1
}
