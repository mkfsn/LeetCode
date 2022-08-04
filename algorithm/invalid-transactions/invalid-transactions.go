package invalid_transactions

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type transaction struct {
	name   string
	minute int
	amount int
	city   string
}

func invalidTransactions(transactions []string) []string {
	decodedTransactions := decodeTransactions(transactions)

	possiblyValidTransactions := make([]*transaction, 0)

	for _, t := range decodedTransactions {
		possiblyValidTransactions = append(possiblyValidTransactions, t)
	}

	sort.Slice(possiblyValidTransactions, func(i, j int) bool {
		if possiblyValidTransactions[i].name == possiblyValidTransactions[j].name {
			return possiblyValidTransactions[i].minute < possiblyValidTransactions[j].minute
		}

		return possiblyValidTransactions[i].name < possiblyValidTransactions[j].name
	})

	invalidTransactions := make([]*transaction, 0)

	for i, cur := range possiblyValidTransactions {
		valid := true

		for j := i - 1; j >= 0; j-- {
			prev := possiblyValidTransactions[j]
			if prev.name != cur.name {
				break
			}

			if cur.minute-prev.minute > 60 {
				break
			}

			if prev.city != cur.city {
				valid = false
				break
			}
		}

		for j := i + 1; j < len(possiblyValidTransactions); j++ {
			next := possiblyValidTransactions[j]
			if next.name != cur.name {
				break
			}

			if next.minute-cur.minute > 60 {
				break
			}

			if next.city != cur.city {
				valid = false
				break
			}
		}

		if cur.amount > 1000 {
			valid = false
		}

		if !valid {
			invalidTransactions = append(invalidTransactions, cur)
		}
	}

	return encodeTransactions(invalidTransactions)
}

func decodeTransactions(inputs []string) []*transaction {
	results := make([]*transaction, len(inputs))

	for i, input := range inputs {
		items := strings.Split(input, ",")

		minute, _ := strconv.Atoi(items[1])
		amount, _ := strconv.Atoi(items[2])

		results[i] = &transaction{
			name:   items[0],
			minute: minute,
			amount: amount,
			city:   items[3],
		}
	}

	return results
}

func encodeTransactions(inputs []*transaction) []string {
	result := make([]string, len(inputs))

	for i, t := range inputs {
		result[i] = fmt.Sprintf("%s,%d,%d,%s", t.name, t.minute, t.amount, t.city)
	}

	return result
}
