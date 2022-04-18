package backspace_string_compare

func backspaceCompare(s string, t string) bool {
	return toWord(s) == toWord(t)
}

func toWord(s string) string {
	buf := make([]byte, 0, len(s))

	for _, c := range s {
		if c == '#' {
			n := len(buf) - 1
			if n < 0 {
				n = 0
			}
			buf = buf[0:n:cap(buf)]
		} else {
			buf = append(buf, byte(c))
		}
	}

	return string(buf)
}
