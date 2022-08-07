package longest_palindrome

func longestPalindrome(s string) int {
	dict := make(map[int32]int)

	for _, ch := range s {
		dict[ch]++
	}

	// count evens and check if there's any odd
	hasOdd := 0
	res := 0

	for _, count := range dict {
		if count%2 == 0 {
			res += count
		} else {
			res += count - 1
			hasOdd = 1
		}
	}

	return res + hasOdd
}
