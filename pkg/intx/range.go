package intx

func InRange(i int, min int, max int) bool {
	return i >= min && i <= max
}

func InRange64(i int64, min int64, max int64) bool {
	return i >= min && i <= max
}
