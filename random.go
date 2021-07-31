package utils

import (
	"math/rand"
	"time"
)

var seed = rand.New(rand.NewSource(time.Now().UnixNano()))

// GenerateRandomNumbers receives an integer x and returns a random number.
// It will use a seed that generates a number from 0 to 99.
// Sum the x + 1 with this random number.
// It returns the a integer random number.
func GenerateRandomNumbers(x int) int {
	return x + seed.Intn(100) + 1
}
