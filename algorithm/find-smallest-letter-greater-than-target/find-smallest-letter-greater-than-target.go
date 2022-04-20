package find_smallest_letter_greater_than_target

import (
	"log"
)

// 22:15

func nextGreatestLetter(letters []byte, target byte) byte {
	low, high := 0, len(letters)-1

	for low < high {
		i := (low + high) / 2
		v := letters[i]
		log.Println(i)

		if v <= target {
			low = i + 1
		} else {
			high = i
		}
	}

	if letters[low] > target {
		return letters[low]
	}

	if low+1 > len(letters)-1 {
		return letters[0]
	}

	return letters[low+1]
}
