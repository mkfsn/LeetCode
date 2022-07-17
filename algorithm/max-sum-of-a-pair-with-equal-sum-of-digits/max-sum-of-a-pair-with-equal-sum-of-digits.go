package max_sum_of_a_pair_with_equal_sum_of_digits

import "fmt"

func maximumSum(nums []int) int {
	digits := make(map[int][]int)

	for _, v := range nums {
		d := calculateDigit(v)
		fmt.Printf("%d -> %d\n", v, d)
		digits[d] = append(digits[d], v)
	}

	max := -1

	for v, pairs := range digits {
		fmt.Printf("%d: %v\n", v, pairs)
		max = getMax(max, maxTwoSum(pairs))
	}

	return max
}

func calculateDigit(v int) int {
	res := 0

	for v > 0 {
		res += v % 10
		v /= 10
	}

	return res
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxTwoSum(arr []int) int {

	if len(arr) < 2 {
		return -1
	} else if len(arr) == 2 {
		return arr[0] + arr[1]
	}

	x, y := arr[0], arr[1]
	if y > x {
		x, y = y, x
	}

	for i := 2; i < len(arr); i++ {
		v := arr[i]

		if v > x {
			x, y = v, x
		} else if v > y {
			y = v
		}
	}

	return x + y
}
