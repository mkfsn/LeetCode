package largest_number_after_digit_swaps_by_parity

import (
	"sort"
)

func largestInteger(num int) int {
	digits := make([]int, 0)
	odds := make([]int, 0)
	evens := make([]int, 0)

	for num > 0 {
		v := num % 10

		digits = append(digits, v)

		if v%2 == 1 {
			odds = append(odds, v)
		} else {
			evens = append(evens, v)
		}

		num /= 10
	}

	sort.Ints(odds)
	sort.Ints(evens)

	n, m := len(odds)-1, len(evens)-1

	for i, j := 0, 0; i < len(odds) || j < len(evens); {
		k := len(digits) - i - j - 1

		v := digits[k]

		if v%2 == 1 {
			digits[k] = odds[n-i]
			i++
		} else {
			digits[k] = evens[m-j]
			j++
		}
	}

	res := 0
	for i := len(digits) - 1; i >= 0; i-- {
		res *= 10
		res += digits[i]
	}

	return res
}
