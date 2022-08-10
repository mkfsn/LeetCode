package binary_trees_with_factors

import (
	"math"
	"sort"
)

func numFactoredBinaryTrees(arr []int) int {
	dict := make(map[int]bool)

	for _, num := range arr {
		dict[num] = true
	}

	counts := make(map[int]int)

	result := len(arr)

	sort.Ints(arr)

	for i, num := range arr {
		// find all element e in arr[:i] and e <= sqrt(num) and dict[num/e] exist

		// v = sqrt(num)
		v := int(math.Sqrt(float64(num)))

		count := 0

		for _, e := range arr[:i] {
			if e > v {
				break
			}

			x := num / e

			if e*x != num {
				continue
			}

			if _, ok := dict[x]; !ok {
				continue
			}

			factor := 1
			if e != x {
				factor++
			}

			countE := counts[e] + 1
			countX := counts[x] + 1

			count += countE * countX * factor
		}

		// store the combinations of num
		counts[num] = count

		result = result + count
	}

	return result % (1000_000_000 + 7)
}

// 2
// 4
// 5
// 10
// 20

// 6.
//  4
// 2 2

// [4] 1

// 7.
//  10
// 2  5

// 8.
//  10
// 5  2

// [10] 2

// 20
// = 2 * 10
//         => 10 [2]
// = 2 * (1+2) = 6
//
// = 4 * 5
//         => 4 = 2 * 2 => [4] 1
// so, # = n * 2, n = {4, (2,2)} = 2 => # = 4

// 9.
//  20
// 4  5

//   20
//  4  5
// 2 2

// 11.
//  20
// 5  4

//  20
// 5  4
//   2 2

//  20
// 2  10

//  20
// 2  10
//   2  5

//  20
// 2  10
//   5  2

//   20
// 10  2

//    20
//  10  2
// 2  5

//    20
//  10  2
// 5  2
