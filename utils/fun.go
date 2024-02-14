package utils

import "strings"

func Sun(n1 int, n2 int) int {
	return n1 + n2
}

func Size(list []int) int {
	if list == nil {
		panic("List value off")
	}
	return len(list)
}

func Contains(value string, value2 string) bool {
	return strings.Contains(value, value2)
}

func ToUpperCase(value string) string {
	return strings.ToUpper(value)
}
