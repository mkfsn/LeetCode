package most_common_word

import (
	"strings"
)

func mostCommonWord(paragraph string, banned []string) string {
	buf := make([]byte, len(paragraph))

	for i, c := range paragraph {
		switch c {
		case ',', '!', '?', '\'', ';', '.':
			buf[i] = ' '
		default:
			buf[i] = toLower(byte(c))
		}
	}

	bannedMap := make(map[string]bool)

	for _, word := range banned {
		bannedMap[word] = true
	}

	countMap := make(map[string]int)

	max := 0
	var res string

	for _, word := range strings.Split(string(buf), " ") {
		if len(word) == 0 {
			continue
		}

		if _, ok := bannedMap[word]; ok {
			continue
		}

		countMap[word]++

		if v := countMap[word]; v > max {
			res = word
			max = v
		}
	}

	return res
}

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c - 'A' + 'a'
	}

	return c
}
