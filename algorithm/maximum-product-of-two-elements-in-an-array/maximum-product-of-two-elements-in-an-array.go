package maximum_product_of_two_elements_in_an_array

func maxProduct(nums []int) int {
	max1 := nums[0]
	max2 := nums[1]

	// make sure max1 > max2
	if max1 < max2 {
		max1, max2 = max2, max1
	}

	for i := 2; i < len(nums); i++ {
		num := nums[i]

		if num < max2 {
			continue
		}

		if num > max1 { // num bigger than both max1 and max2
			max1, max2 = num, max1
		} else {
			max2 = num
		}
	}

	return (max1 - 1) * (max2 - 1)
}
