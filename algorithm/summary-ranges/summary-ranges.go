package summary_ranges

import (
	"fmt"
)

type Range struct {
	from *int
	cur  *int
	to   *int
}

func (r Range) String() string {
	if r.to == nil {
		return fmt.Sprintf("%d", *r.from)
	}

	return fmt.Sprintf("%d->%d", *r.from, *r.to)
}

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	ranges := make([]Range, 1)
	cur := &ranges[0]

	for _, num := range nums {
		num := num

		if cur.from == nil {
			cur.from, cur.cur = &num, &num
			continue
		}

		if num-*cur.cur == 1 {
			cur.cur, cur.to = &num, &num
			continue
		}

		ranges = append(ranges, Range{})
		cur = &ranges[len(ranges)-1]

		cur.from, cur.cur = &num, &num
	}

	var res []string

	for _, r := range ranges {
		res = append(res, r.String())
	}

	return res
}
