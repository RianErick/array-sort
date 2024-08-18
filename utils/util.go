package utils

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"

	"github.com/RianErick/go-learning/game"
)

type Matrix struct {
	Array []int
	Real  bool
}

func ClearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func DisplayAndWait(array []int, duration time.Duration) {
	fmt.Printf("You have %v seconds to memorize it, run: %v\n", duration.Seconds(), array)
	for i := 0; i < int(duration.Seconds()); i++ {
		time.Sleep(time.Second)
		fmt.Printf("===")
	}
	ClearConsole()
	fmt.Println(game.Title())
}

func Display() {

	var duration time.Duration = 10 * time.Second

	fmt.Printf("You have %v seconds to memorize it\n", duration.Seconds())
	for i := 0; i < int(duration.Seconds()); i++ {
		time.Sleep(time.Second)
		fmt.Printf("===")
	}
	ClearConsole()
	fmt.Println(game.Title())
}

func GenerateRandomArray(size int, max int) []int {
	array := make([]int, size)
	for i := 0; i < size; i++ {
		array[i] = rand.Intn(max)
	}
	return array
}

func GenerateMultipleRandomArrays(count, size, max int) [][]int {
	arrays := make([][]int, count)
	for i := 0; i < count; i++ {
		arrays[i] = GenerateRandomArray(size, max)
	}
	return arrays
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
