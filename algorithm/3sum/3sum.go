package _sum

import (
	"sort"
)

// find i, j, k such that
// i != j != k and
// nums[i] + nums[j] + nums[k] == 0
// ->
// -nums[i] == nums[j] nums[k]

func threeSum(nums []int) [][]int {
	sort.Ints(nums) // O(nlogn)

	res := make([][]int, 0)
	seen := make(map[int]bool) // O(n)

	for i := 0; i < len(nums)-2; i++ { // O(n * O(twoSum))
		if nums[i] > 0 { // early return, as it's not possible in the sorted array
			break
		}

		if _, ok := seen[nums[i]]; ok {
			continue
		}
		seen[nums[i]] = true

		// find -nums[i] in nums[i+1:]
		for _, pair := range twoSum(nums[i+1:], -nums[i]) {
			res = append(res, append([]int{nums[i]}, pair...))
		}
	}

	return res
}

// nums is sorted
func twoSum(nums []int, target int) [][]int {
	dict := make(map[int]int)
	for i, v := range nums {
		dict[v] = i
	}

	seen := make(map[int]bool)
	res := make([][]int, 0)

	for i := 0; i < len(nums)-1; i++ {
		v := nums[i]

		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = true

		x := target - v // v < x

		if v > x { // early return, as it's not possible in the sorted array
			break
		}

		if idx, ok := dict[x]; ok && i < idx {
			res = append(res, []int{v, x})
		}
	}

	return res
}
