package main

import (
	"fmt"
	"reflect"
)

func main() {
	nome := "Gabriel"
	idade := 26
	temperatura := 36.7

	fmt.Println("Ola", nome, "voce possui", idade, "anos, hoje esta", temperatura, "graus")
	fmt.Println("O tipo da variavel é ", reflect.TypeOf(temperatura))

	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do programa")

	var comando int

	fmt.Println("Digite a opção desejada: ")
	fmt.Scan(&comando)

	if comando == 1 {
		fmt.Println("1 - Iniciar monitoramento")
	} else if comando == 2 {
		fmt.Println("2 - Exibir Logs")
	} else if comando == 0 {
		fmt.Println("0 - Sair do programa")
	} else {
		fmt.Println("Opção inválida")
	}
}
