package sum_of_all_subset_xor_totals

func subsetXORSum(nums []int) int {
	numOfSubsets := 1 << len(nums)
	result := 0

	for subset := 1; subset < numOfSubsets; subset++ {
		curXOR := 0

		for i, bits := 0, subset; i < len(nums); i, bits = i+1, bits>>1 {
			if bits&1 != 0 {
				curXOR ^= nums[i]
			}
		}

		result += curXOR
	}

	return result
}

// 11:38
// 11:55
// 17min
