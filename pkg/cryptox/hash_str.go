package cryptox

func SHA256Str(s string) string {
	return SHA256([]byte(s))
}

func SHA512Str(s string) string {
	return SHA512([]byte(s))
}

func SHA384Str(s string) string {
	return SHA384([]byte(s))
}

func SHA512_224Str(s string) string {
	return SHA512_224([]byte(s))
}

func SHA512_256Str(s string) string {
	return SHA512_256([]byte(s))
}
