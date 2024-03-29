package main

import (
	"fmt"
	"math/rand"
	"time"

	"os"
	"os/exec"

	"github.com/RianErick/go-learning/game"
	"github.com/RianErick/go-learning/utils"
)

func clearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	fmt.Println(game.Title())

	var arrayUser = make([]int, 10)
	var arrayMemory0 = make([]int, 10)
	var arrayMemory1 = make([]int, 10)

	rand.Seed(int64(time.Now().Nanosecond()))

	for i := 0; i < len(arrayUser); i++ {
		arrayUser[i] = rand.Intn(100)
	}

	fmt.Printf("You have 10 seconds to memorize it, run: %v\n", arrayUser)

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Printf("===")
	}
	clearConsole()
	fmt.Println(game.Title())

	for i := 0; i < 10; i++ {
		arrayMemory0[i] = rand.Intn(100)
		for j := 0; j < 10; j++ {
			arrayMemory1[i] = rand.Intn(100)
		}
	}

	arrayTrue := utils.SortList(arrayUser)
	i0 := utils.SortList(arrayMemory0)
	i1 := utils.SortList(arrayMemory1)

	matrixOfArray := []utils.Matrix{
		{Array: arrayTrue, Real: true},
		{Array: i0, Real: false},
		{Array: i1, Real: false},
	}
	var selectOption int

	fmt.Println("Select option true:")
	for i := 0; i < len(matrixOfArray); i++ {
		fmt.Printf("%v : %v\n", i, matrixOfArray[i].Array)
	}
	fmt.Scanf("%d", &selectOption)

	if matrixOfArray[selectOption].Real {
		fmt.Println("Inner :)")
	} else {
		fmt.Println("F :0")
	}

}
