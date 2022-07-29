package minimum_amount_of_time_to_fill_cups

func fillCups(amount []int) int {
	a, b, c := amount[0], amount[1], amount[2]

	max := getMax(getMax(a, b), c)
	sum := a + b + c

	if sum <= 2*max {
		return max
	}

	res := sum / 2

	if sum%2 == 1 {
		res++
	}

	return res
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 5 4 4 <- 0
// 4 3 4 , 1
// 3 3 3 , 2
// 2 2 3 , 3
// 1 2 2 , 4
// 1 1 1 , 5
// 0 0 1 , 6
// 0 0 0 , 7

// 6 4 4, 0
// 5 3 4, 1
// 4 3 3, 2
// 3 2 3, 3
// 2 2 2, 4
// 1 1 2, 5
// 0 1 1, 6
// 0 0 0, 7

// 6 4 3, 0
// 5 3 3
// 4 2 3
// 3 2 2
// 2 1 2
// 1 1 1
// 0 0 1
// 0 0 0, 7

// 6 5 5, 0
// 5 4 5
// 4 4 4
// 3 3 4
// 2 3 3
// 1 2 3
// 0 2 2
// 0 1 1
// 0 0 0, 8

// 3 2 3
// 2 1 2
// 1 1 1
// 0 0 1
// 0 0 0, 4

// assume
// a, b, c
// find x, y, z where x >= y >= z
// if x >= y + z: // if x+y+z <= 2x
//    return x
// else:
//    return x + max(y-x1, z-x2); while y-x1 > 0 and z-x2 > 0 && x1+x2 == x
