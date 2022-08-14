package count_sorted_vowel_strings

func countVowelStrings(n int) int {
	// 1: 1+1+1+1+1 => 5 (5*1)
	// 2: (1+1+1+1+1)+(1+1+1+1)+(1+1+1)+(1+1)+1 = (1+5)*5/2
	//  : 5+4+3+2+1 = 15 (5*3)
	// 3: (5+4+3+2+1)+(4+3+2+1)+(3+2+1)+(2+1)+1 = 35 (5*7)
	//  : 15+10+6+3+1
	// 4: (15+10+6+3+1)+(10+6+3+1)+(6+3+1)+(3+1)+1 = 70 (5*14)
	//  : 35+20+10+4+1
	//  : f(3) + (f(3)-f(2)) + ? + ? + 1
	// 5: (35+20+10+4+1)+(20+10+4+1)+(10+4+1)+(4+1)+1 = 126
	//  : f(4) + (f(4)-f(3)) + ? + ? + 1

	// 1  + 1  + 1  + 1 + 1
	// 5  + 4  + 3  + 2 + 1
	// 15 + 10 + 6  + 3 + 1
	// 35 + 20 + 10 + 4 + 1
	// 70 + 35 + 15 + 5 + 1

	permutations := [5]int{ // O(5) => O(1)
		1, 1, 1, 1, 1,
	}

	for i := 0; i < n; i++ { // O(n)
		for j := len(permutations) - 2; j >= 0; j-- { // O(4) => O(1)
			permutations[j] += permutations[j+1]
		}
	}

	// Time: O(n) * O(4) => O(n)
	// Space: O(5) => O(1)

	return permutations[0]
}