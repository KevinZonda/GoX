package intx

func Curve(val, min, max int) int {
	if val < min {
		return min
	}
	if val > max {
		return max
	}
	return val
}

func CurveMax(val, max int) int {
	if val > max {
		return max
	}
	return val
}

func CurveMin(val, min int) int {
	if val < min {
		return min
	}
	return val
}
