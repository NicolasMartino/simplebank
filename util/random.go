package util

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Generates a random number between min and max
func RandomInt(min, max int64) int64 {
	d := max - min
	return min + rand.Int63n(d+1)
}

// Generates a random string of length n
func RandomString(n int) string {
	if n < 0 {
		panic("RandomString: a string length cannot be under 0")
	}
	var sb strings.Builder
	alphabetSize := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(alphabetSize)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// Generates a random owner name
func RandomName() string {
	return RandomString(6)
}

//Generates a random email
func RandomEmail() string {
	return fmt.Sprintf("%s@%s.%s", RandomString(8), RandomString(5), RandomString(3))
}

// Generates a random money amount between 0 and 100000
func RandomMoney() float64 {
	return RoundToTwoDigits(rand.Float64() * math.Pow10(int(rand.Int63n(9))))
}

//Generates a random currency
func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "AUD"}
	return currencies[rand.Intn(len(currencies))]
}
