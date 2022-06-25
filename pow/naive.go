package pow

import "math"

func abs(a float64) float64 {
	if a < 0 {
		return -a
	}
	return a
}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if x == 0 || x == 1 {
		return x
	}

	res := math.Exp(float64(n) * math.Log(abs(x)))

	if x < 0 && (n-1)%2 == 0 {
		return -res
	}

	return res
}
