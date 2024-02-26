package main

import (
	"fmt"
	"math/rand"
	"time"

	"os"
	"os/exec"

	"github.com/RianErick/go-learning/player"
	"github.com/RianErick/go-learning/utils"
)

func clearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	fmt.Println(player.Title())

	var arrayUser = make([]int, 10)
	var arrayMemory0 = make([]int, 10)
	var arrayMemory1 = make([]int, 10)

	rand.Seed(int64(time.Now().Nanosecond()))

	for i := 0; i < len(arrayUser); i++ { // Cria um array de 10 posicoes randomico
		arrayUser[i] = rand.Intn(100)
	}

	fmt.Printf("You have 10 seconds to memorize it, run: %v\n", arrayUser)

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second) // Tempo para o jogado memorizar
		fmt.Printf("===")
	}
	clearConsole()
	fmt.Println(player.Title())

	for i := 0; i < 10; i++ {
		arrayMemory0[i] = rand.Intn(100) //Cria as outras opcoes de escolha
		for j := 0; j < 10; j++ {
			arrayMemory1[i] = rand.Intn(100)
		}
	}

	arrayTrue := utils.SortList(arrayUser) // Ordena o array
	i0 := utils.SortList(arrayMemory0)
	i1 := utils.SortList(arrayMemory1)

	matrixOfArray := []utils.Matrix{
		{Array: arrayTrue, Real: true},
		{Array: i0, Real: false}, //Matrix de arrays
		{Array: i1, Real: false},
	}
	var selectOption int

	fmt.Println("Select option true:")
	for i := 0; i < len(matrixOfArray); i++ {
		fmt.Printf("%v : %v\n", i, matrixOfArray[i].Array)
	}
	fmt.Scanf("%d", &selectOption)

	fmt.Println(selectOption)

	selectIndex := utils.FilterIsReal(matrixOfArray, matrixOfArray[selectOption].Real)

	if selectIndex[selectOption].Real {
		fmt.Println("Ganhou")
	} else {
		fmt.Println("Perdeu")
	}
}
