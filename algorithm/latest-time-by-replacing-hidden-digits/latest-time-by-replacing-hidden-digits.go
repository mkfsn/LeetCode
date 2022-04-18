package latest_time_by_replacing_hidden_digits

import (
	"fmt"
)

func maximumTime(time string) string {
	return latestHour(time[:2]) + ":" + latestMinute(time[3:])
}

func latestHour(s string) string {
	switch {
	case s == "??":
		return "23"

	case s[0] == '?': // ?1, ?2, ?9
		if s[1] <= '3' {
			return fmt.Sprintf("2%c", s[1])
		} else {
			return fmt.Sprintf("1%c", s[1])
		}

	case s[1] == '?': // 0?, 1?, 2?
		if s[0] == '2' {
			return "23"
		} else {
			return fmt.Sprintf("%c9", s[0])
		}
	}

	return s
}

func latestMinute(s string) string {
	switch {
	case s == "??":
		return "59"

	case s[0] == '?': // ?0, ?1, ?2, ..., ?9
		return fmt.Sprintf("5%c", s[1])

	case s[1] == '?': // 0?, 1?, 2?, ..., 5?
		return fmt.Sprintf("%c9", s[0])

	}

	return s
}
