package check_if_n_and_its_double_exist

func checkIfExist(arr []int) bool {
	table := make(map[int]int)

	for j, v := range arr {
		if v%2 == 0 {
			table[v/2] = j
		}
	}

	for i, v := range arr {
		if j, ok := table[v]; ok && i != j {
			return true
		}
	}

	return false
}
