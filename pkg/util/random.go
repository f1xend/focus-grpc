package util

import (
	"math/rand"
	"time"
)

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}
