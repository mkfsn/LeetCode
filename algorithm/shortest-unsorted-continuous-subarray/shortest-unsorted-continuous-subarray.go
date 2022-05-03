package shortest_unsorted_continuous_subarray

import (
	"fmt"
)

func findUnsortedSubarray(nums []int) int {
	i, j := 0, len(nums)-1

	for i < len(nums)-1 {
		if !lessThan(nums[i], nums[i+1:]) {
			break
		}
		i++
	}

	for j > 0 {
		if !greaterThan(nums[j], nums[:j]) {
			break
		}
		j--
	}

	fmt.Println(i, j)

	if j <= i {
		return 0
	}

	return j - i + 1
}

func greaterThan(v int, arr []int) bool {
	for _, n := range arr {
		if n > v {
			return false
		}
	}
	return true
}

func lessThan(v int, arr []int) bool {
	for _, n := range arr {
		if n < v {
			return false
		}
	}
	return true
}
