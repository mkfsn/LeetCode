package search_suggestions_system

import (
	"sort"
	"strings"
)

func suggestedProducts(products []string, searchWord string) [][]string {
	results := make([][]string, len(searchWord))

	sort.Strings(products)

	for i := 0; i < len(searchWord); i++ {
		prefix := searchWord[:i+1]

		cnt := 0
		for _, product := range products {
			if cnt >= 3 {
				break
			}
			if strings.HasPrefix(product, prefix) {
				results[i] = append(results[i], product)
				cnt++
			}
		}
	}

	return results
}
