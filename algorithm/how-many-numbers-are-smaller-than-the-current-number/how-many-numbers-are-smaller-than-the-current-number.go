package how_many_numbers_are_smaller_than_the_current_number

func smallerNumbersThanCurrent(nums []int) []int {
	counts := make(map[int]int)
	for _, num := range nums {
		counts[num] = 0
	}

	for _, num := range nums {
		for k := range counts {
			if num < k {
				counts[k]++
			}
		}
	}

	res := make([]int, len(nums))

	for i, v := range nums {
		res[i] = counts[v]
	}

	return res
}
