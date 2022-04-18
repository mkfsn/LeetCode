package x_of_a_kind_in_a_deck_of_cards

func hasGroupsSizeX(deck []int) bool {
	dict := make(map[int]int)

	for _, n := range deck {
		dict[n]++
	}

	maxCount := 0

	for _, v := range dict {
		if v < 2 {
			return false
		}

		if v > maxCount {
			maxCount = v
		}
	}

	for i := 2; i <= maxCount; i++ {
		valid := true

		for _, v := range dict {
			if v%i != 0 {
				valid = false
				break
			}
		}

		if valid {
			return true
		}
	}

	return false
}
