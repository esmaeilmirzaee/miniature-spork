package helpers

import (
	"math/rand"
	"time"
)

func GenerateRandomNumber(upperBoundLimit int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(upperBoundLimit)
}