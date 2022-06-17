package countingbits

func countBits(n int) []int {
	if n == 0 {
		return []int{0}
	}
	bits_arr := make([]int, n+1)
	bits_arr[1] = 1
	for i := 2; i <= n; i++ {
		bits_arr[i] = bits_arr[i>>1] + i&1
	}
	return bits_arr
}
