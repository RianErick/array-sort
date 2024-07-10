package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"

	"github.com/RianErick/go-learning/game"
	"github.com/RianErick/go-learning/utils"
)

func clearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func generateRandomArray(size int, max int) []int {
	array := make([]int, size)
	for i := 0; i < size; i++ {
		array[i] = rand.Intn(max)
	}
	return array
}

func displayAndWait(array []int, duration time.Duration) {
	fmt.Printf("You have %v seconds to memorize it, run: %v\n", duration.Seconds(), array)
	for i := 0; i < int(duration.Seconds()); i++ {
		time.Sleep(time.Second)
		fmt.Printf("===")
	}
	clearConsole()
	fmt.Println(game.Title())
}

func generateMultipleRandomArrays(count, size, max int) [][]int {
	arrays := make([][]int, count)
	for i := 0; i < count; i++ {
		arrays[i] = generateRandomArray(size, max)
	}
	return arrays
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(game.Title())

	arrayUser := generateRandomArray(10, 100)
	displayAndWait(arrayUser, 10*time.Second)

	arrayOptions := generateMultipleRandomArrays(5, 10, 100) // Generate 5 random options
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
	fmt.Scanf("%d", &selectOption)

	if matrixOfArray[selectOption].Real {
		fmt.Println("Inner :)")
	} else {
		fmt.Println("F :0")
	}
}
