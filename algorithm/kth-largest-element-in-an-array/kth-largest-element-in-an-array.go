package kth_largest_element_in_an_array

func findKthLargest(nums []int, k int) int {
	// fmt.Printf("[k=%d] %v\n", k, nums)
	pivotIndex := findPivot(nums)
	targetIndex := len(nums) - k

	// fmt.Printf("[pivotIndex=%d/targetIndex=%d] %v\n", pivotIndex, targetIndex, nums)

	if pivotIndex == targetIndex {
		return nums[pivotIndex]
	} else if pivotIndex < targetIndex {
		// len(nums) - pivotIndex - (targetIndex - pivotIndex)
		return findKthLargest(nums[pivotIndex+1:], len(nums)-targetIndex)
	} else { // pivotIndex > targetIndex
		return findKthLargest(nums[:pivotIndex], pivotIndex-targetIndex)
	}
}

func findPivot(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	pivotValue := nums[0]

	i, j := 0, len(nums)-1

	for i < j {
		for i < len(nums)-1 && nums[i] <= pivotValue {
			i++
		}

		for j > 0 && nums[j] >= pivotValue {
			j--
		}

		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	if i > j {
		if nums[0] > nums[j] {
			nums[0], nums[j] = nums[j], nums[0]
		}
		return j
	}

	if nums[0] > nums[i] {
		nums[0], nums[i] = nums[i], nums[0]
	}
	return i
}

// 9:20 start
