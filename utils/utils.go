package utils

import (
	"os"
	"regexp"
)

// ========================
// FILES IO
// ========================
func ReadFile(title string) string {
	data, err := os.ReadFile(title)
	if err != nil {
		panic("Error reading file")
	}
	return string(data)
}

// ========================
// STRINGS
// ========================
func QuickMatch(str, pattern string) []string {
	re := regexp.MustCompile(pattern)
	return re.FindAllString(str, -1)
}

// ========================
// MATH
// ========================
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	if a < 0 {
		return -a // keep result positive
	}
	return a
}

func LCM(a, b int) int {
	return a / GCD(a, b) * b
}
