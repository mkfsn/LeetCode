package longest_string_chain

import (
	"sort"
)

func longestStrChain(words []string) int {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	// w1/w2
	//      "" a b ba  bca  bda  bdca
	// ""   0  1 1 0   0    0    0
	// a    0  0 0 1+a 0    0    0
	// b    0  0 0 1+b 0    0    0
	// ba   0  0 0 0   1+ba 1+ba 0
	// bca  0  0 0 0   0    0    1+bca
	// bda                       1+bda
	// bdca

	pages := make(map[int][]string)

	for _, word := range words {
		l := len(word)
		pages[l] = append(pages[l], word)
	}

	chainLengths := make(map[string]int)
	result := 0

	for _, cur := range words {
		length := 0

		for _, prev := range pages[len(cur)-1] {
			if !isPredecessor(prev, cur) {
				continue
			}

			if l := chainLengths[prev]; l > length {
				length = l
			}
		}

		chainLengths[cur] = length + 1

		if chainLengths[cur] > result {
			result = chainLengths[cur]
		}
	}

	return result
}

// Return true if w1 is predecessor of w2
func isPredecessor(w1, w2 string) bool {
	if len(w1) != len(w2)-1 {
		return false
	}

	// possibilities are
	// 1. w1 + x == w2 (x append to w1)
	// 2. x + w1 == w2 (x prepend to w1)
	// 3. w(+x)1 == w2 (x inset to w1)

	if w1 == w2[:len(w2)-1] { // 1
		return true
	} else if w1 == w2[1:] { // 2
		return true
	}

	firstDiffIdx := 0

	for ; firstDiffIdx < len(w1); firstDiffIdx++ {
		if w1[firstDiffIdx] != w2[firstDiffIdx] {
			break
		}
	}

	return w1[firstDiffIdx:] == w2[firstDiffIdx+1:]
}
