package util

import (
	"math"
	"math/rand"
	"time"
)

func RandInt() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(math.MaxInt)
}
