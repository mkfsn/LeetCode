package longer_contiguous_segments_of_ones_than_zeros

import (
	"fmt"
)

func checkZeroOnes(s string) bool {
	maxOnes, maxZeros := 0, 0
	curOnes, curZeros := 0, 0

	for _, c := range s {
		if c == '0' {
			if curOnes > maxOnes {
				maxOnes = curOnes
			}
			curOnes = 0
			curZeros++
		} else if c == '1' {
			if curZeros > maxZeros {
				maxZeros = curZeros
			}

			curZeros = 0
			curOnes++
		}
	}

	if curOnes > maxOnes {
		maxOnes = curOnes
	}

	if curZeros > maxZeros {
		maxZeros = curZeros
	}

	fmt.Println(maxOnes, maxZeros)

	if maxOnes > maxZeros {
		return true
	}

	return false
}

// 23:38
