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

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do programa...")
	default:
		fmt.Println("Opção inválida")
	}
}
