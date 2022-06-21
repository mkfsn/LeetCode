package n_th_tribonacci_number

// T0 = 0, T1 = 1, T2 = 1, and Tn+3 = Tn + Tn+1 + Tn+2 for n >= 0.
func tribonacci(n int) int {
	size := n + 1
	if size <= 2 {
		size = 2
	}

	nums := make([]int, size)

	nums[0] = 0
	nums[1] = 1
	nums[2] = 1

	for i := 3; i <= n; i++ {
		nums[i] = nums[i-3] + nums[i-2] + nums[i-1]
	}

	return nums[n]
}
