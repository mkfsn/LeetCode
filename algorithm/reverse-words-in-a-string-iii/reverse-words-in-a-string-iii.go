package reverse_words_in_a_string_iii

func reverseWords(s string) string {
	arr := []byte(s)

	for i := 0; i < len(arr); {
		for i < len(arr) && arr[i] == ' ' {
			i++
		}

		j := i + 1
		for j < len(arr) && arr[j] != ' ' {
			j++
		}

		reverseWord(arr[i:j])

		i = j + 1
	}

	return string(arr)
}

func reverseWord(arr []byte) {
	i, j := 0, len(arr)-1
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}

// 17:24 start
// 17:32 first submit
