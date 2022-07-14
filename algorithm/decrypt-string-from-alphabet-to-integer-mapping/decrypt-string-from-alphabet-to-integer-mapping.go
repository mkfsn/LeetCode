package decrypt_string_from_alphabet_to_integer_mapping

func freqAlphabets(s string) string {
	res := make([]byte, 0, len(s))

	for i := 0; i < len(s); i++ {
		if i+2 < len(s) && s[i+2] == '#' {
			res = append(res, decode2(s[i:i+2]))
			i += 2
		} else {
			res = append(res, decode1(s[i]))
		}
	}

	return string(res)
}

func decode1(b byte) byte {
	return b - '1' + 'a'
}

func decode2(s string) byte {
	return (s[0]-'0')*10 + (s[1] - '0') - 10 + 'j'
}
