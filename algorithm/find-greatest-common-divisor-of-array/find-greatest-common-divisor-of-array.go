package find_greatest_common_divisor_of_array

func findGCD(nums []int) int {
	min, max := nums[0], nums[0]

	for _, num := range nums {
		if num < min {
			min = num
		} else if num > max {
			max = num
		}
	}

	return gcd(min, max)
}

func gcd(a, b int) int {
	// | a  | b |
	// | 12 | 6 |
	// |----|
	// | 12 |
	// |  0 |

	// a, b = b, a % b
	// if b == 0 {
	//     return a
	// }

	if a < b {
		a, b = b, a
	}

	for b != 0 {
		a, b = b, a%b
	}

	return a
}
