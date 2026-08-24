package util

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

const alphabets = "abcdefghijklmnopqrstuvwxyz"

// RandomInt generates random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int64N(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabets)

	for i := 0; i < n; i++ {
		c := alphabets[rand.IntN(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency generates a random currency code
func RandomCurrency() string {
	currencies := []string{USD, EUR, CAD}
	n := len(currencies)
	return currencies[rand.IntN(n)]
}

// RandomEmail generates a new random email
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}
