package decode_xored_array

func decode(encoded []int, first int) []int {
	arr := make([]int, len(encoded)+1)
	arr[0] = first

	// encoded[i] = arr[i] XOR arr[i + 1]
	// encoded[i-1] = arr[i-1] XOR arr[i]
	// encoded[i-1] XOR arr[i-1] = arr[i]

	for i := 1; i < len(arr); i++ {
		arr[i] = arr[i-1] ^ encoded[i-1]
	}

	return arr
}

// 21:07
// 21:12 submit 1
