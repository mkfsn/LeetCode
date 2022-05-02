package reverse_string_ii

func reverseStr(s string, k int) string {
	arr := []byte(s)

	for i := 0; i < len(arr); i += 2 * k {
		j := i + k

		for j > len(arr) {
			j--
		}

		reverseArr(arr[i:j])
	}

	return string(arr)
}

func reverseArr(arr []byte) {
	// log.Printf("(%d) %s\n", len(arr), arr)

	i, j := 0, len(arr)-1

	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}

// 17:42 submit: 1
// 17:45 submit: 2
// 17:54 submit: 3
