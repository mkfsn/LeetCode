package maximum_product_difference_between_two_pairs

import (
	"sort"
)

func maxProductDifference(nums []int) int {
	sort.Ints(nums)
	n := len(nums)

	return nums[n-1]*nums[n-2] - nums[0]*nums[1]
}
