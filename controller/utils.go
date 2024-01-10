package hangmanweb

import (
	"math/rand"
	"strings"
	"time"
)

func ValidChars(seq string) bool {
	for _, char := range seq {
		if char < 97 || seq[0] > 122 {
			return false
		}
	}
	return true
}

func ValidFileName(seq string) bool {
	if len(seq) <= 1 {
		return false
	}

	if len(strings.Split(seq, "/")) > 1 {
		return false
	}

	if len(seq) >= 3 && seq[len(seq)-3:] == ".go" {
		return false
	}

	return true
}

func Lower(seq string) string {
	result := ""
	for _, char := range seq {
		if rune(char) >= 65 && rune(char) <= 90 {
			result += string(rune(char) + 32)
		} else {
			result += string(rune(char))
		}
	}
	return result
}

func RandInt(max int) int {
	return rand.Intn(max)
}

func RandInit() {
	rand.Seed(time.Now().UnixNano())
}

func IsIn(l string, seq []string) bool {
	for _, char := range seq {
		if char == l {
			return true
		}
	}
	return false
}

func IsInInt(i int, seq []int) bool {
	for _, char := range seq {
		if char == i {
			return true
		}
	}
	return false
}

func IsInString(l string, seq string) bool {
	for _, char := range seq {
		if string(char) == l {
			return true
		}
	}
	return false
}

func RemoveAccents(seq string) string {
	noAccent := ""
	for _, char := range seq {
		isAccent := false
		for k, v := range Accents {
			if IsInString(string(char), v) {
				noAccent += k
				isAccent = true
				break
			}
		}
		if !isAccent {
			noAccent += string(char)
		}
	}
	return noAccent
}
