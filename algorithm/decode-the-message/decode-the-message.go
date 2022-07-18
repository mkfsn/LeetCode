package decode_the_message

func decodeMessage(key string, message string) string {
	mapper := make(map[int32]int32)
	idx := int32(0)

	for _, c := range key {
		if c == ' ' {
			continue
		}

		if _, ok := mapper[c]; ok {
			continue
		}

		mapper[c] = 'a' + idx
		idx++
	}

	var res string

	for _, c := range message {
		if c == ' ' {
			res += " "
		} else {
			res += string(mapper[c])
		}
	}

	return res
}

// 10:37
// 10:51
// 14min
