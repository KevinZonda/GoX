package intx

import "strconv"

func ChangeBase(i int64, base int) string {
	return strconv.FormatInt(i, base)
}

// SplitInGroups splits total into groupCount groups.
// total: 5, groupCount: 3 => [2, 2, 1]
// total: 5, groupCount: 2 => [3, 2]
func SplitInGroups(total, groupCount int) []int {
	g := total / groupCount
	groups := make([]int, groupCount)
	for i := 0; i < groupCount; i++ {
		groups[i] = g
	}
	for i := 0; i < total%groupCount; i++ {
		groups[i]++
	}
	return groups
}
