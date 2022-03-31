package main

import "fmt"

func main() {
	nome, idade, _ := devolverNomeEIdade()
	fmt.Println("Ola", nome, "voce possui", idade, "anos")
}

func devolverNomeEIdade() (string, int, string) {
	nome := "Gabriel"
	idade := 26
	naoInteressa := "Não interessa"
	return nome, idade, naoInteressa
}
