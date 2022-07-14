package find_positive_integer_solution_for_a_given_equation

func findSolution(customFunction func(int, int) int, z int) [][]int {
	res := make([][]int, 0)

	y := 1000

	for x := 1; x <= 1000; x++ {
		for y > 1 && customFunction(x, y) > z {
			y--
		}

		if customFunction(x, y) == z {
			res = append(res, []int{x, y})
		}
	}

	return res
}
