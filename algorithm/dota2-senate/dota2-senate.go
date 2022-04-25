package dota2_senate

// n == senate.length
// 1 <= n <= 104
// senate[i] is either 'R' or 'D'.
func predictPartyVictory(senate string) string {
	senators := make([]bool, len(senate))

	countR, countD := 0, 0

	for i, c := range senate {
		senators[i] = true

		if c == 'R' {
			countR++
		} else if c == 'D' {
			countD++
		}
	}

	next := func(c byte, cur int) (int, bool) {
		iter := len(senate)

		for iter >= 0 {
			cur++

			if n := len(senate); cur >= n {
				cur = cur % n
			}

			if senate[cur] == c && senators[cur] == true {
				break
			}

			iter--
		}

		if iter < 0 {
			return -1, false
		}

		return cur, true
	}

	iter := 0
	for countR != 0 && countD != 0 {
		for i, c := range senate {
			if senators[i] == false { // banned
				continue
			}

			switch c {
			case 'R':
				idx, ok := next('D', i)
				// log.Printf("[%d][%c] %d %v %d:%d\n", iter, c, idx, ok, countR, countD)
				if !ok {
					break
				}
				senators[idx] = false
				countD--
			case 'D':
				idx, ok := next('R', i)
				// log.Printf("[%d][%c] %d %v %d:%d\n", iter, c, idx, ok, countR, countD)
				if !ok {
					break
				}
				senators[idx] = false
				countR--
			}
		}

		iter++
	}

	if countR > countD {
		return "Radiant"
	}

	return "Dire"
}

// overtime
// 65min
