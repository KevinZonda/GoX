package boolx

func Any(vs []bool) bool {
	for _, v := range vs {
		if v {
			return true
		}
	}
	return false
}

func All(vs []bool) bool {
	for _, v := range vs {
		if !v {
			return false
		}
	}
	return true
}
