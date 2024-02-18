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

	type ArraySort struct {
		stream []int
		check  bool
	}

	var arrayUser = make([]int, 10)

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
	fmt.Println(player.Title())

	arrayTrue := utils.SortList(arrayUser)

	fmt.Println(arrayTrue)

}
