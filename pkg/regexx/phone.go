package regexx

import "regexp"

var ChPhoneRegexp = regexp.MustCompile("^1(?:3[0-9]|4[5-9]|5[0-9]|6[12456]|7[0-8]|8[0-9]|9[0-9])[0-9]{8}$")
