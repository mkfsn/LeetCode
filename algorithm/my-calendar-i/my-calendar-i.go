package my_calendar_i

type MyCalendar struct {
	events [][]int
}

func Constructor() MyCalendar {
	return MyCalendar{
		events: make([][]int, 0),
	}
}

//     start-----end
//      |-------|
//    |-----|
//  e[0]   e[1]

//    start-----end
//      |-------|
//           |-----|
//          e[0]  e[1]

//    start-----end
//      |--------|
//       |-----|
//      e[0]  e[1]

//    start-----end
//      |-------|
//    |-----------|
//   e[0]        e[1]

//    start-----end
//      |-------|
//        |-----------|
//       e[0]        e[1]

//      start-----end
//        |-------|
//  |-----------|
//  e[0]        e[1]

func (this *MyCalendar) Book(start int, end int) bool {
	for _, e := range this.events {
		if e[0] <= start && start < e[1] {
			return false
		}

		if e[0] < end && end <= e[1] {
			return false
		}

		if start <= e[0] && e[0] < end {
			return false
		}

		if start < e[1] && e[1] <= end {
			return false
		}
	}

	this.events = append(this.events, []int{start, end})

	return true
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
