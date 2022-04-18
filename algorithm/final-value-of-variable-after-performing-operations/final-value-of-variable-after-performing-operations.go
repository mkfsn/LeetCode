package final_value_of_variable_after_performing_operations

type Operation string

const (
	XPlusPlus   Operation = "X++"
	PlusPlusX             = "++X"
	XMinusMinus           = "X--"
	MinusMinusX           = "--X"
)

func finalValueAfterOperations(operations []string) int {
	x := 0
	for _, op := range operations {

		switch Operation(op) {
		case XPlusPlus, PlusPlusX:
			x++
		case XMinusMinus, MinusMinusX:
			x--
		}

	}

	return x
}
