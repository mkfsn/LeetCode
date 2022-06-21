package increasing_triplet_subsequence

func increasingTriplet(nums []int) bool {
	max := nums[0]

	for _, n := range nums {
		if max < n {
			max = n
		}
	}

	c1, c2 := max, max

	for _, n := range nums {
		if c1 >= n {
			c1 = n
		} else if c2 >= n {
			c2 = n
		} else {
			return true
		}
	}

	return false
}
