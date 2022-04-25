package maximum_score_after_splitting_a_string

// 2 <= s.length <= 500
// The string s consists of characters '0' and '1' only.
func maxScore(s string) int {
	// Init
	max := 0

	if s[0] == '0' {
		max++
	}

	for i := len(s) - 1; i >= 1; i-- {
		if s[i] == '1' {
			max++
		}
	}

	cur := max

	// Iterate: Starting from 1 ~ len(s) - 2
	for i := 1; i < len(s)-1; i++ {
		if s[i] == '0' {
			cur++
		} else if s[i] == '1' {
			cur--
		}

		if cur > max {
			max = cur
		}
	}

	return max
}

// 15min
