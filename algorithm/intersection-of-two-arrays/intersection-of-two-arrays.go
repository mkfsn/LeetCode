package intersection_of_two_arrays

func intersection(nums1 []int, nums2 []int) []int {
	seen := make(map[int]int8)
	for _, v := range nums1 {
		seen[v] |= 1
	}

	for _, v := range nums2 {
		seen[v] |= 2
	}

	res := make([]int, 0)

	for v, cnt := range seen {
		if cnt == 3 {
			res = append(res, v)
		}
	}

	return res
}
