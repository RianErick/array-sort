package utils

import (
	"fmt"
	"strings"
)

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

func For(lista []string) {

	for _, item := range lista {
		fmt.Println(item)
	}
}

type Person struct {
	name string
}

type Matrix struct {
	Array []int
	Real  bool
}

func SortList(sort []int) []int {
	size := len(sort)
	for i := 0; i < size; i++ {

		for j := i + 1; j < size; j++ {

			if sort[j] < sort[i] {
				temp := sort[j]
				sort[j] = sort[i]
				sort[i] = temp
			}
		}
	}
	return sort
}

func FilterIsReal(matrixOfArray []Matrix, option bool) []Matrix {

	var filtro []Matrix

	for _, p := range matrixOfArray {
		if p.Real == option {
			filtro = append(filtro, p)
		}

	}
	return filtro
}

// func Shuffle(slice []Matrix) {
// 	rand.Shuffle(len(slice), func(i, j int) {
// 		slice[i], slice[j] = slice[j], slice[i]
// 	})
// }
