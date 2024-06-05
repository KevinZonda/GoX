package stringx

import "strconv"

func Int32(s string) (int32, error) {
	i, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return -1, err
	}
	return int32(i), err
}

func Int32Or(s string, or int32) int32 {
	i, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return or
	}
	return int32(i)
}

func ToInt32Force(s string) int32 {
	i, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		panic(err)
	}
	return int32(i)
}

func Int(s string) (int, error) {
	return strconv.Atoi(s)
}

func IntOr(s string, or int) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return or
	}
	return i
}

func ToIntForce(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func Int64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

func Int64Or(s string, or int64) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return or
	}
	return i
}

func ToInt64Force(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}

func Float64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

func Float64Or(s string, or float64) float64 {
	i, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return or
	}
	return i
}

func ToFloat64Force(s string) float64 {
	i, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return i
}

func Bool(s string) (bool, error) {
	return strconv.ParseBool(s)
}

func BoolOr(s string, or bool) bool {
	i, err := strconv.ParseBool(s)
	if err != nil {
		return or
	}
	return i
}

func ToBoolForce(s string) bool {
	i, err := strconv.ParseBool(s)
	if err != nil {
		panic(err)
	}
	return i
}

func Ptr(s string) *string {
	return &s
}

func AutoPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
