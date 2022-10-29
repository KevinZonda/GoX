package stringx

func Contains(ss []string, c string) (index int, isContains bool) {
	if ss == nil {
		return -1, false
	}

	for i, s := range ss {
		if s == c {
			return i, true
		}
	}
	return -1, false
}
