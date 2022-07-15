package power_of_four

func isPowerOfFour(n int) bool {
	if n <= 0 {
		return false
	} else if n == 1 {
		return true
	} else if n%4 != 0 {
		return false
	}

	return isPowerOfFour(n / 4)
}
