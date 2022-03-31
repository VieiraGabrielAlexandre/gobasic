package main

import "fmt"

func main() {
	lista()
}

func lista() {
	lista := []string{"Pão", "Leite", "Carne", "Queijo", "Açucar", "Café", "Cerveja"}

	for i := 0; i < len(lista); i++ {
		fmt.Println(lista[i])
	}

	for i, lista := range lista {
		fmt.Println("Item:", i, "Nome:", lista)
	}
}
