package strings

import (
	"math/rand"
)

func Rand(source string, length uint) string {
	letterRunes := []rune(source)
	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
