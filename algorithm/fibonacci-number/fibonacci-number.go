package fibonacci_number

func fib(n int) int {
	cache := map[int]int{
		0: 0,
		1: 1,
		2: 1,
	}

	return doFib(n, cache)
}

func doFib(n int, cache map[int]int) int {
	if v, ok := cache[n]; ok {
		return v
	}

	v := doFib(n-1, cache) + doFib(n-2, cache)
	cache[n] = v

	return v
}
