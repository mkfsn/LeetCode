package relative_sort_array

func relativeSortArray(arr1 []int, arr2 []int) []int {
	dict := make(map[int]int)

	for _, v := range arr2 {
		dict[v] = 0
	}

	notAppearElements := make([]int, 0, len(arr1))

	for _, v := range arr1 {
		count, ok := dict[v]
		if !ok {
			notAppearElements = append(notAppearElements, v)
			continue
		}
		dict[v] = count + 1
	}

	res := make([]int, 0, len(arr1))

	for _, v := range arr2 {
		for count := dict[v]; count > 0; count-- {
			res = append(res, v)
		}
	}

	// fmt.Println("Before", notAppearElements)
	quickSort(notAppearElements)
	// fmt.Println("After", notAppearElements)

	for _, v := range notAppearElements {
		res = append(res, v)
	}

	return res
}

func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	pivotIndex := findPivotIndex(arr)
	quickSort(arr[:pivotIndex])
	quickSort(arr[pivotIndex+1:])
}

func findPivotIndex(arr []int) int {
	pivotVal := arr[0]
	left, right := 1, len(arr)-1

	for left < right {

		for left < len(arr) && arr[left] <= pivotVal {
			left++
		}

		for right > 0 && arr[right] >= pivotVal {
			right--
		}

		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}

	// fmt.Printf("left: %d, right: %d, arr: %v\n", left, right, arr)

	if left < right {
		if arr[left] < arr[0] {
			arr[left], arr[0] = arr[0], arr[left]
		}
		return left
	}

	if arr[right] < arr[0] {
		arr[right], arr[0] = arr[0], arr[right]
	}
	return right
}

// 20:33
// 21:04 submit 1
