package number_of_valid_words_in_a_sentence

import (
	"strings"
)

func countValidWords(sentence string) int {
	count := 0

	for _, word := range strings.Split(sentence, " ") {
		if isValidWord(word) {
			count++
			// log.Printf("%s: valid\n", word)
		} else {
			// log.Printf("%s: invalid\n", word)
		}
	}

	return count
}

func isValidWord(word string) bool {
	if len(word) == 0 {
		return false
	}

	if word[0] == '-' {
		return false
	}

	if word[len(word)-1] == '-' {
		return false
	}

	hyphensCount := 0

	for i, c := range word {
		switch c {
		case '!', '.', ',':
			if i != len(word)-1 {
				return false
			}
		case '-':
			hyphensCount++

			if !isSurroundedByLowercaseCharacters(word, i) {
				return false
			}

		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			return false
		}

		if hyphensCount > 1 {
			return false
		}
	}

	return true
}

func isSurroundedByLowercaseCharacters(word string, i int) bool {
	return i > 0 && i < len(word)-1 && isLowercaseCharacters(word[i-1]) && isLowercaseCharacters(word[i+1])
}

func isLowercaseCharacters(c uint8) bool {
	return c >= 'a' && c <= 'z'
}
