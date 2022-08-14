package final_prices_with_a_special_discount_in_a_shop

type Stack struct {
	items []interface{}
}

func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() interface{} {
	n := s.Len()
	last := s.items[n-1]
	s.items = s.items[:n-1]
	return last
}

func (s Stack) Len() int {
	return len(s.items)
}

func (s Stack) Top() interface{} {
	return s.items[s.Len()-1]
}

func finalPrices(prices []int) []int {
	discounts := make([]int, len(prices)) // O(n)
	var st Stack                          // O(n)

	for i := len(prices) - 1; i >= 0; i-- { // O(n)
		for st.Len() > 0 && st.Top().(int) > prices[i] {
			st.Pop()
		}

		if st.Len() == 0 {
			discounts[i] = -1
		} else {
			discounts[i] = st.Top().(int)
		}

		st.Push(prices[i])
	}

	// x   [-> nil]
	// -1, [-> 3]
	// -1, [-> 2]
	// 2, [-> 6 -> 2]
	// 2, [-> 4 -> 2]
	// 4, [-> 2]

	results := make([]int, len(discounts)) // O(n)

	for i, price := range prices { // O(n)
		if discounts[i] == -1 {
			results[i] = price
		} else {
			results[i] = price - discounts[i]
		}
	}

	return results
}
