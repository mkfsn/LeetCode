package find_first_palindromic_string_in_the_array

func firstPalindrome(words []string) string {
	for _, word := range words {
		if isPalindromic(word) {
			return word
		}
	}

	return ""
}

func isPalindromic(word string) bool {
	l, r := 0, len(word)-1

	for l < r {
		if word[l] != word[r] {
			return false
		}
		l++
		r--
	}

	return true
}

// 11:25
// 11:26
