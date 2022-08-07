package decode_string

type pair struct {
	text string
	num  int
}

func decodeString(s string) string {
	// 3[a2[c]]
	// stack
	// top -> nil
	// detect number: 3
	// detect '[', push: 3 and ""
	// top -> 3 -> ""
	// detect text: a
	// detect number: 2
	// detect '[', push: 2 and "a"
	// top -> 2 -> a -> 3 -> ""
	// detect text: c
	// detect ']', pop 2 and a
	// top -> 3 -> ""
	// cur = a + 2 * c = acc
	// detect ']', pop 3 and ""
	// cur = "" + 3 * acc
	// cur = accaccacc

	stack := make([]pair, 0, len(s)/2)
	curText := ""
	curNum := 0

	for _, ch := range s {
		switch {
		case ch == '[':
			stack = append(stack, pair{num: curNum, text: curText})
			curText, curNum = "", 0
		case ch == ']':
			pre := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			for i := 0; i < pre.num; i++ {
				pre.text += curText
			}
			curText = pre.text
		case ch >= '0' && ch <= '9':
			curNum = curNum*10 + int(ch-'0')
		default:
			curText += string(ch)
		}
	}

	return curText
}
