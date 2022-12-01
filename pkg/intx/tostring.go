package intx

import "strconv"

func ToString(i int) string {
	return strconv.Itoa(i)
}

func ToString64(i int64) string {
	return strconv.FormatInt(i, 10)
}

func ToString32(i int32) string {
	return ToString(int(32))
}