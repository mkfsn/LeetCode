package longest_palindromic_substring

func longestPalindrome(s string) string {
	res := longestPalindromeSlice([]byte(s))
	return string(res)
}

func longestPalindromeSlice(s []byte) []byte {
	if len(s) <= 1 {
		return s
	}

	start, end := 0, 0

	for i := 0; i < len(s)-1; i++ {
		len1 := findLargerPalindromeSlice(s, i, i)
		len2 := findLargerPalindromeSlice(s, i, i+1)

		size := getMax(len1, len2)
		if size <= end-start {
			continue
		}

		if len1 > len2 {
			start = i - size/2
			end = i + size/2 + 1
		} else {
			start = i - (size-1)/2
			end = i + size/2 + 1
		}
	}

	return s[start:end]
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findLargerPalindromeSlice(s []byte, left int, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}
