package number_of_segments_in_a_string

func countSegments(s string) int {
	cnt := 0

	for i := 0; i < len(s); i++ {
		for i < len(s) && s[i] == ' ' {
			i++
		}

		if i >= len(s) {
			break
		}

		cnt++

		for i < len(s) && s[i] != ' ' {
			i++
		}
	}

	return cnt
}
