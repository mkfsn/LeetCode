package unique_morse_code_words

var morseCodes = []string{
	".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.",
	"--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--..",
}

func uniqueMorseRepresentations(words []string) int {
	seen := make(map[string]bool)

	for _, word := range words {
		var code string

		for _, ch := range word {
			code += morseCodes[ch-'a']
		}

		seen[code] = true
	}

	return len(seen)
}
