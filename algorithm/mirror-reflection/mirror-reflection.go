package mirror_reflection

// 5, 1
// 0, 2
// 5, 3
// 0, 4
// 5, 5
// k == 5
// x == 5
// y == 5
// return 1

// 5, 2
// 0, 4
// 5, 6
// 0, 8
// 5, 10
// k == 5
// x == 5
// y == 10
// return 0

// 4, 1
// 0, 2
// 4, 3
// 0, 4
// k == 4
// x == 0
// y == 4

func mirrorReflection(p int, q int) int {
	// find k where q * k % p == 0

	// if k % 2 == 0: return 2
	// if y % 2 == 0: return 0
	// return 1

	for p%2 == 0 && q%2 == 0 {
		p /= 2
		q /= 2
	}

	k := 1
	for q*k%p != 0 {
		k++
	}

	if k%2 == 0 {
		return 2
	}

	if (k*q)%2 == 0 {
		return 0
	}

	return 1
}
