package binary_watch

import (
	"fmt"
)

func readBinaryWatch(turnedOn int) []string {
	res := make([]string, 0)

	for i := 0; i <= turnedOn; i++ {
		x, y := i, turnedOn-i

		hours, ok := readBinaryWatchHour(x)
		if !ok {
			continue
		}

		minutes, ok := readBinaryWatchMinute(y)
		if !ok {
			continue
		}

		// fmt.Printf("(%d,%d) hours=%v, minutes=%v\n", x, y, hours, minutes)

		for _, hour := range hours {
			for _, minute := range minutes {
				res = append(res, fmt.Sprintf("%d:%02d", hour, minute))
			}
		}
	}

	return res
}

// 1, 2, 4, 8, 16, 32
func readBinaryWatchHour(turnedOn int) ([]int, bool) {
	switch turnedOn {
	case 0:
		return []int{0}, true
	case 1:
		return []int{1, 2, 4, 8}, true
	case 2:
		return []int{3, 5, 6, 9, 10}, true
	case 3:
		return []int{7, 11}, true
	}

	return nil, false
}

func readBinaryWatchMinute(turnedOn int) ([]int, bool) {
	if turnedOn >= 6 {
		return nil, false
	} else if turnedOn == 0 {
		return []int{0}, true
	}

	return permutation([]int{1, 2, 4, 8, 16, 32}, turnedOn), true
}

func permutation(list []int, count int) []int {
	if count > len(list) {
		return nil
	} else if count == 1 {
		return list
	}

	res := make([]int, 0, len(list))

	for i := 0; i < len(list); i++ {
		arr := permutation(list[i+1:], count-1)
		for _, v := range arr {
			num := list[i] + v
			if num > 59 {
				break
			}
			res = append(res, num)
		}
	}

	return res
}
