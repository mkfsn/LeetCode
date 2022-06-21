package is_subsequence

import (
	"fmt"
)

func isSubsequence(s string, t string) bool {
	return isSubsequenceSlice([]byte(s), []byte(t))
}

func isSubsequenceSlice(s []byte, t []byte) bool {
	fmt.Printf("%s, %s\n", s, t)
	if len(s) != 0 && len(t) == 0 {
		return false
	} else if len(s) == 0 {
		return true
	}

	x := s[0]

	for i, v := range t {
		if v == x {
			return isSubsequenceSlice(s[1:], t[i+1:])
		}
	}

	return false
}
