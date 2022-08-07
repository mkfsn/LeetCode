package longest_repeating_character_replacement

func characterReplacement(s string, k int) int {
	// find all [start,end] that tolerate k operations to make all characters the same
	//
	// Create two pointers: start and end.
	//
	// The `end` pointer starts at 0, try to find the right most position
	//
	// if end - start - maxSoFar + 1 > k, we move the start pointer, and update the counter
	//
	//
	start, end := 0, 0
	counter := make(map[uint8]int)
	maxLength := 0
	result := 0

	for ; end < len(s); end++ {
		counter[s[end]]++
		// if s[end] becomes the characters that occurs the most in the current window
		// update the maxLength
		if v := counter[s[end]]; v > maxLength {
			maxLength = v
		}

		// end-start+1: the length of the window
		// maxLength: the maximum number of the character that occurs the most in the current window
		// end-start+1-maxLength: needed replacement count
		for end-start+1-maxLength > k {
			// move the start pointer so that:
			// end - start - maxLength + 1 == k
			// which makes the current window [start,end] to have exact k operations
			counter[s[start]]--
			start++
		}

		// update the result
		if v := end - start + 1; v > result {
			result = v
		}
	}

	return result
}
