package intx

import "strconv"

func ChangeBase(i int64, base int) string {
	return strconv.FormatInt(i, base)
}
