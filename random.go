package dax

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func Rand(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}
