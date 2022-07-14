package sort_an_array

func sortArray(nums []int) []int {
	return quickSort(nums)
}

func quickSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	left, right := 0, len(nums)-1

	for i := range nums {
		if nums[i] < nums[right] {
			nums[i], nums[left] = nums[left], nums[i]
			left++
		}
	}

	nums[left], nums[right] = nums[right], nums[left]

	quickSort(nums[:left])
	quickSort(nums[left+1:])

	return nums
}

// func mergeSort(nums []int) []int {
// 	if len(nums) < 2 {
// 		return nums
// 	}
//
// 	mid := len(nums) / 2
//
// 	return merge(mergeSort(nums[:mid]), mergeSort(nums[mid:]))
// }
//
// func merge(left []int, right []int) []int {
// 	size := len(left) + len(right)
// 	i, j := 0, 0
//
// 	res := make([]int, size)
//
// 	for k := 0; k < size; k++ {
// 		if i >= len(left) && j < len(right) {
// 			res[k] = right[j]
// 			j++
// 		} else if i < len(left) && j >= len(right) {
// 			res[k] = left[i]
// 			i++
// 		} else if left[i] < right[j] {
// 			res[k] = left[i]
// 			i++
// 		} else {
// 			res[k] = right[j]
// 			j++
// 		}
// 	}
//
// 	return res
// }
