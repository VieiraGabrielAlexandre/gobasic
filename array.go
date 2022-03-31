package main

import (
	"fmt"
	"reflect"
)

func main() {

	nomes()
}

func nomes() {
	nomes := []string{"Gabriel", "Alexandre", "Vieira"}
	fmt.Println(nomes)
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println("Tamanho: ", len(nomes), "Capacidade: ", cap(nomes))

	nomes = append(nomes, "Juliana")
	fmt.Println(nomes)
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println("Tamanho: ", len(nomes), "Capacidade: ", cap(nomes))
}
