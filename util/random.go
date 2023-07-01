package util

import (
	"math/rand"
	"strings"
	"time"
)

const (
	alphabet                      = "qwertyuiopasdfghjklzxcvbnm"
	alphabetWithNumbers           = "qwertyuiopasdfghjklzxcvbnm1234567890"
	alphabetWithNumbersAndSymbols = "qwertyuiopasdfghjklzxcvbnm1234567890|][{}':;?/><.,!№;%:?*()_+=-@#$^&` "
	alphabetForEmailName          = "qwertyuiopasdfghjklzxcvbnm1234567890-_."
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomStrWithAlphabet(n int, alphabet string) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomStr generates a random string with length n
func RandomStr(n int) string {
	return randomStrWithAlphabet(n, alphabet)
}

// RandomEmail generates a random email
func RandomEmail() string {
	return randomStrWithAlphabet(12, alphabetForEmailName) + "@" + randomStrWithAlphabet(4, alphabet) + "." + randomStrWithAlphabet(3, alphabet)
}

// RandomPassword generates a random password
func RandomPassword(n int) string {
	return randomStrWithAlphabet(n, alphabetWithNumbers)
}

// CoinFlip return 0 or 1
func CoinFlip() int64 {
	return rand.Int63n(2)
}

// RandomBool return random boolean value
func RandomBool() bool {
	return rand.Int63n(2) != 0
}
