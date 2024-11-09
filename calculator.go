// calculator.go
package calculator

import (
	"math/rand"
)

// Cal calculates the sum of val1, val2, and a random integer between 0 and 5.
func Cal1(val1 int, val2 int) int {
	val3 := rand.Intn(6)
	return val1 + val2 + val3
}

func Sayaword(name string) string {
	return name
}
