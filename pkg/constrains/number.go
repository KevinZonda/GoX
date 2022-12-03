package constrains

type Int interface {
	int8 | int16 | int32 | int64
}

type Uint interface {
	uint8 | uint16 | uint32 | uint64
}

type Float interface {
	float32 | float64
}

type Number interface {
	Int | Uint | Float
}
