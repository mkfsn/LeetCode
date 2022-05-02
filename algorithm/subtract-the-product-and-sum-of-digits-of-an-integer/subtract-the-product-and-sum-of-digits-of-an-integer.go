package subtract_the_product_and_sum_of_digits_of_an_integer

func subtractProductAndSum(n int) int {
	product := 1
	sum := 0

	for n > 0 {
		v := n % 10
		product *= v
		sum += v
		n /= 10
	}

	return product - sum
}

// 00:49
