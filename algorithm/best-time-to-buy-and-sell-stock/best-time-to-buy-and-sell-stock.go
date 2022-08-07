package best_time_to_buy_and_sell_stock

func maxProfit(prices []int) int {
	min := prices[0]
	res := 0

	for i := 1; i < len(prices); i++ {
		v := prices[i]

		if v-min > res {
			res = v - min
		}

		if v < min {
			min = v
		}
	}

	return res
}
