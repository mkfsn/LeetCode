package greatest_common_divisor_of_strings

func gcdOfStrings(str1 string, str2 string) string {
	l1, l2 := len(str1), len(str2)
	if l1 < l2 {
		l1, l2 = l2, l1
		str1, str2 = str2, str1
	}

	for i := len(str2); i >= 1; i-- {
		t := str2[:i]

		// check if str1 = t+...+t
		// check if str2 = t+..+t

		if l1%i != 0 || l2%i != 0 {
			continue
		}

		tmp2 := ""

		for j := 0; j < l2/i; j++ {
			tmp2 += t
		}

		if tmp2 != str2 {
			continue
		}

		tmp1 := ""

		for j := 0; j < l1/i; j++ {
			tmp1 += t
		}

		if tmp1 != str1 {
			continue
		}

		return t
	}

	return ""
}
