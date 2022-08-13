package special_array_with_x_elements_greater_than_or_equal_x

func specialArray(nums []int) int {
	// [3, 5]
	// 1 => want 1 but 2
	// 2 => want 2 and 2 -> 2

	// [0, 0]
	// 1 => want 1 but 0
	// 0 => want 0 but 2

	// [0, 4, 3, 0, 4]
	// [0, 0, 3, 4, 4]
	// 2 => want 2 but 3
	// 3 => want 3 and 3
	// --
	// 4 => want 4 but 2

	left, right := 0, len(nums)

	for left+1 < right { // util left is next to right
		mid := (left + right) / 2
		cnt := countGreaterThanOrEqualTo(nums, mid)

		if mid == cnt {
			return mid
		} else if mid < cnt { // -> right
			left = mid
		} else { // left <-
			right = mid
		}
	}

	if cnt := countGreaterThanOrEqualTo(nums, left); cnt == left {
		return left
	}

	if cnt := countGreaterThanOrEqualTo(nums, right); cnt == right {
		return right
	}

	return -1
}

func countGreaterThanOrEqualTo(nums []int, target int) int { // O(n)
	cnt := 0

	for _, n := range nums {
		if n >= target {
			cnt++
		}
	}

	return cnt
}
