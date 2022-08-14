package goal_parser_interpretation

func interpret(command string) string {
	var result string

	for i := 0; i < len(command); i++ {
		switch command[i] {
		case 'G':
			result += "G"
		case '(':

			if command[i+1] == ')' {
				result += "o"
				i++
			} else {
				result += "al"
				i += 3
			}
		}

	}

	return result
}
