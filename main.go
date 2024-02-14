package main

import (
	"fmt"

	"github.com/RianErick/go-learning/utils"
)

func main() {
	fmt.Println("Hello World")

	soma := utils.Sun(1, 2)

	fmt.Println(soma)

	lista := []int{0, 1, 2, 1, 5, 9}

	fmt.Printf("lista[0]: %v\n", lista[3])

	listaDeString := []string{"K"}

	if len(listaDeString) == 0 {
		panic("Lista Vazia! :p")
	}

	type Person struct {
		name string
	}

	var p Person
	p.name = "Neto"

	fmt.Printf("p.name: %v\n", p.name)
	var listaBool bool = false

	if len(lista) == 6 {
		listaBool = true
		fmt.Println(listaBool)
	}

	for _, item := range lista {
		fmt.Println(item)
	}

	meuNomeMinusculo := "rian erick"

	meuNomeMaiusculo := utils.ToUpperCase(meuNomeMinusculo)

	fmt.Println(meuNomeMaiusculo)

}
