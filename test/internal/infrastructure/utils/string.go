package utils

import (
	"math/rand"
	"strings"
	"time"
)

var (
	chars = []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
		"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
	}
)

// RandString 生成随机字符串(a-zA-Z0-9)
func RandString(length int) string {
	rand.Seed(time.Now().UnixNano())
	str := strings.Builder{}

	for i := 0; i < length; i++ {
		str.WriteString(chars[rand.Intn(62)])
	}

	return str.String()
}

// RandDigit 生成随机字符串(0-9)
func RandDigit(length int) string {
	rand.Seed(time.Now().UnixNano())
	str := strings.Builder{}

	for i := 0; i < length; i++ {
		str.WriteString(chars[52+rand.Intn(10)])
	}

	return str.String()
}

// RandCharacter 生成随机字符串(a-zA-Z)
func RandCharacter(length int) string {
	rand.Seed(time.Now().UnixNano())
	str := strings.Builder{}

	for i := 0; i < length; i++ {
		str.WriteString(chars[rand.Intn(52)])
	}

	return str.String()
}
