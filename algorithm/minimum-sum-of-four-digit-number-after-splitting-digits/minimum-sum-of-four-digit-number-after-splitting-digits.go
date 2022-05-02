package minimum_sum_of_four_digit_number_after_splitting_digits

func minimumSum(num int) int {
	arr := make([]int8, 4)
	arr[0] = int8(num / 1000)
	arr[1] = int8((num / 100) % 10)
	arr[2] = int8((num / 10) % 10)
	arr[3] = int8(num % 10)

	// sort arr: make sure arr[0], arr[1] < arr[2], arr[3]
	quickSort(arr)

	return int(arr[0]*10) + int(arr[3]) + int(arr[1]*10) + int(arr[2])
}

func quickSort(arr []int8) {
	// 9876

	// 8967

	if arr[3] < arr[0] {
		arr[3], arr[0] = arr[0], arr[3]
	}

	// arr[0] < arr[3]

	if arr[2] < arr[1] {
		arr[2], arr[1] = arr[1], arr[2]
	}

	// arr[1] < arr[2]

	if arr[3] < arr[1] {
		arr[1], arr[3] = arr[3], arr[1]
	}

	if arr[2] < arr[0] {
		arr[0], arr[2] = arr[2], arr[0]
	}
}

// 00:23
