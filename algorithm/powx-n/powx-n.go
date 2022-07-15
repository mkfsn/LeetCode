package powx_n

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n == 1 {
		return x
	} else if n == -1 {
		return 1 / x
	}

	if n%2 == 0 {
		half := myPow(x, n/2)
		return half * half
	} else {
		return x * myPow(x, n-1)
	}
}
