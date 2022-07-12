package get_maximum_in_generated_array

func getMaximumGenerated(n int) int {
	arr := make([]int, n+1)

	arr[0] = 0
	res := 0

	if n > 0 {
		arr[1] = 1
		res = 1
	}

	for i := 2; i <= n; i++ {
		if i%2 == 0 {
			arr[i] = arr[i/2]
		} else {
			x := i / 2
			arr[i] = arr[x] + arr[x+1]
		}

		res = getMax(res, arr[i])
	}

	return res
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
