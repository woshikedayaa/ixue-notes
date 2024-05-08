package utils

import (
	"math/rand/v2"
	"strings"
)

var charsets = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

func RandomString(length int) string {
	builder := strings.Builder{}
	builder.Grow(length)
	for i := 0; i < length; i++ {
		builder.WriteByte(charsets[rand.Int()%len(charsets)])
	}
	return builder.String()
}
