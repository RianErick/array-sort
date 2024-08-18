package main

import (
	"fmt"
	"math/rand"

	"time"

	"github.com/RianErick/go-learning/game"
	"github.com/RianErick/go-learning/utils"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(game.Title())

	arrayUser := utils.GenerateRandomArray(30, 100)
	utils.DisplayAndWait(arrayUser, 10*time.Second)

	arrayOptions := utils.GenerateMultipleRandomArrays(8, 30, 100)
	arrayTrue := utils.SortList(arrayUser)
	matrixOfArray := []utils.Matrix{
		{Array: arrayTrue, Real: true},
	}

	for _, arr := range arrayOptions {
		sortedArr := utils.SortList(arr)
		matrixOfArray = append(matrixOfArray, utils.Matrix{Array: sortedArr, Real: false})
	}

	rand.Shuffle(len(matrixOfArray), func(i, j int) {
		matrixOfArray[i], matrixOfArray[j] = matrixOfArray[j], matrixOfArray[i]
	})

	var selectOption int
	fmt.Println("Select option true:")

	for i, matrix := range matrixOfArray {
		fmt.Printf("%v : %v\n", i, matrix.Array)
	}

	utils.Display()

	fmt.Scanf("%d", &selectOption)
	fmt.Println("Select option true:")

	if matrixOfArray[selectOption].Real {
		fmt.Println("Inner :)")
	}

	fmt.Println("You faild :(")
}
