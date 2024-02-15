package main

import (
	"fmt"

	"github.com/RianErick/go-learning/player"
)

func main() {
	fmt.Println(player.Title())

	type ArraySort struct {
		stream []int
		check  bool
	}

	var oneArray ArraySort
	var twoArray ArraySort

	oneArray.stream = []int{1, 4, 9, 7, 3, 2, 0}
	oneArray.check = false

	twoArray.stream = []int{1, 2, 4, 52, 2, 3, 0}
	twoArray.check = true

	//Logica para fazer o usuario escolher qual do array seria a escolha correta
}
