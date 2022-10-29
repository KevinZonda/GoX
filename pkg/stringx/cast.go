package stringx

import "strconv"

func Int32(s string) (int32, error) {
	i, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return -1, err
	}
	return int32(i), err
}

func Int(s string) (int, error) {
	return strconv.Atoi(s)
}

func Int64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

func Bool(s string) (bool, error) {
	return strconv.ParseBool(s)
}

func Ptr(s string) *string {
	return &s
}
