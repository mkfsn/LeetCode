package next_greater_element_i

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	dict := make(map[int]int)

	for idx, num := range nums2 {
		dict[num] = idx
	}

	res := make([]int, len(nums1))

	for i, num := range nums1 {

		res[i] = findGreaterElement(num, nums2[dict[num]:])
	}

	return res
}

func findGreaterElement(target int, nums []int) int {
	for _, num := range nums {
		if num > target {
			return num
		}
	}
	return -1
}
