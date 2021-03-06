package utils

import (
	"github.com/91go/faker/data"
	"math/rand"
	"time"
)

// Check if in lib
func dataCheck(dataVal []string) bool {
	var checkOk bool

	_, checkOk = data.Data[dataVal[0]]
	if len(dataVal) == 2 && checkOk {
		_, checkOk = data.Data[dataVal[0]][dataVal[1]]
	}

	return checkOk
}

// Get Random Value
func GetRandValue(dataVal []string) string {
	if !dataCheck(dataVal) {
		return ""
	}
	rand.Seed(time.Now().UnixNano())
	list := data.Data[dataVal[0]][dataVal[1]]
	return list[rand.Intn(len(list))]
}

// Replace # with numbers
func ReplaceWithNumbers(str string) string {
	if str == "" {
		return str
	}
	bytestr := []byte(str)
	hashtag := []byte("#")[0]
	numbers := []byte("0123456789")
	for i := 0; i < len(bytestr); i++ {
		if bytestr[i] == hashtag {
			bytestr[i] = numbers[rand.Intn(9)]
		}
	}
	if bytestr[0] == []byte("0")[0] {
		bytestr[0] = numbers[rand.Intn(8)+1]
	}

	return string(bytestr)
}

// Replace ? with letters
func ReplaceWithLetters(str string) string {
	if str == "" {
		return str
	}
	bytestr := []byte(str)
	question := []byte("?")[0]
	letters := []byte("abcdefghijklmnopqrstuvwxyz")
	for i := 0; i < len(bytestr); i++ {
		if bytestr[i] == question {
			bytestr[i] = letters[rand.Intn(26)]
		}
	}

	return string(bytestr)
}

// Generate random letter
func RandLetter() string {
	return string([]byte("abcdefghijklmnopqrstuvwxyz")[rand.Intn(26)])
}

// Generate random integer between min and max
func RandIntRange(min, max int) int {
	if min == max {
		return min
	}
	return rand.Intn((max+1)-min) + min
}

// Generate random floating-point between min and max
func randFloatRange(min, max float64) float64 {
	if min == max {
		return min
	}
	return rand.Float64()*(max-min) + min
}
