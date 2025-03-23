package randx

import "math/rand"

func RndRange(min, max int) int {
	return min + rand.Intn(max-min)
}

func RndRangeFP64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func RndRangeFP32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}
