package arrayx

func NotEmpty(arr []string) []string {
	if len(arr) == 0 {
		return arr
	}
	return Where(arr, func(s string) bool {
		return s != ""
	})
}
