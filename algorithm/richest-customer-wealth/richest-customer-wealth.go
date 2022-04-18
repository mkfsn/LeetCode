package richest_customer_wealth

func maximumWealth(accounts [][]int) int {
	max := 0
	for _, account := range accounts {
		total := 0
		for _, num := range account {
			total += num
		}

		if total > max {
			max = total
		}
	}
	return max
}
