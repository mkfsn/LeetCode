package valid_boomerang

const (
	max = 2<<32 - 1
)

func isBoomerang(points [][]int) bool {
	var samePoint bool

	slopeValues := make([]float64, 3)

	slopeValues[0], samePoint = calculateSlope(points[0], points[1])
	if samePoint {
		return false
	}

	slopeValues[1], samePoint = calculateSlope(points[1], points[2])
	if samePoint {
		return false
	}

	slopeValues[2], samePoint = calculateSlope(points[0], points[2])
	if samePoint {
		return false
	}

	if slopeValues[0] == slopeValues[1] || slopeValues[1] == slopeValues[2] || slopeValues[0] == slopeValues[2] {
		return false
	}

	return true
}

func calculateSlope(p1 []int, p2 []int) (float64, bool) {
	if p2[1] == p1[1] && p2[0] == p1[0] {
		return 0, true
	}

	y := float64(p2[1] - p1[1])
	x := float64(p2[0] - p1[0])

	if x == 0 {
		return max, false
	}

	res := y / x

	return res, false
}
