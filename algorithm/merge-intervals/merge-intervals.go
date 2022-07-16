package merge_intervals

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	// sort intervals by start
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	res := make([][]int, 0, 1)
	res = append(res, intervals[0])

	for i := 1; i < len(intervals); i++ {
		interval := intervals[i]
		lastRes := res[len(res)-1]

		start, stop := interval[0], interval[1]
		_, lastStop := lastRes[0], lastRes[1]

		// check if (start, stop) overlaps (lastStart, lastStop)
		//     lastStart        lastStop
		//     |----------------|
		//                     |--------------|
		//                     start          stop
		// or
		//     lastStart                        lastStop
		//     |--------------------------------|
		//                     |--------------|
		//                     start          stop

		if start <= lastStop {
			lastRes[1] = getMax(lastStop, stop)
		} else {
			res = append(res, interval)
		}
	}

	return res
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
