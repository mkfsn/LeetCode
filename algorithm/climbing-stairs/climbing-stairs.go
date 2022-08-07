package climbing_stairs

func climbStairs(n int) int {
	cache := map[int]int{
		0: 0,
		1: 1,
		2: 2,
	}
	return doClimbStairs(n, cache)
}

func doClimbStairs(n int, cache map[int]int) int {
	if v, ok := cache[n]; ok {
		return v
	}

	res := doClimbStairs(n-1, cache) + doClimbStairs(n-2, cache)
	cache[n] = res
	return res
}
