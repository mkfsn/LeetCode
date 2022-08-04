package cinema_seat_allocation

import (
	"fmt"
)

func maxNumberOfFamilies(n int, reservedSeats [][]int) int {

	// max possible number : 2 * n

	// for each reserved seat
	//      row, col := seat[0], seat[1]
	//      if col == 1 || col == 10 { continue }
	//
	//      record := records[row]
	//      switch col {
	//      case 2, 3:
	//          record ^= 1
	//      case 4, 5:
	//          record ^= 3
	//      case 6, 7:
	//          record ^= 6
	//      case 8, 9:
	//          record ^= 4
	//      }

	excluded := make(map[int]int)
	for _, seat := range reservedSeats {
		row, col := seat[0], seat[1]
		if col == 1 || col == 10 {
			continue
		}

		record := excluded[row]

		switch col {
		case 2, 3:
			record |= 1
		case 4, 5:
			record |= 3
		case 6, 7:
			record |= 6
		case 8, 9:
			record |= 4
		}

		excluded[row] = record
	}

	count := 2 * n

	for row, record := range excluded {

		fmt.Printf("row[%d]: %d\n", row, record)

		switch record {
		case 1, 3, 4, 5, 6:
			count--
		case 7:
			count -= 2
		}
	}

	return count
}
