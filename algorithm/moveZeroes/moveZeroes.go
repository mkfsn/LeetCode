package moveZeroes

func moveZeroes(nums []int) {
	// two pointers
	zeros, others, n := 0, 0, len(nums)

	for others < n {
		for zeros < n && nums[zeros] != 0 {
			zeros++
		}

		for others < n && nums[others] == 0 {
			others++
		}

		if others >= n {
			break
		}

		if zeros < others {
			// swap
			nums[zeros], nums[others] = nums[others], nums[zeros]

			// reset zeros and others
			others = zeros + 1
			zeros = others
			//
		} else {
			// find another "others"
			others = zeros + 1
		}
	}
}
