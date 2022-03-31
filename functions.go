package main

import "fmt"

func main() {
	nome := capturaNome()
	fmt.Println("Ola", nome)

	peso := capturarPeso()
	altura := capturarAltura()

	imc := calcularImc(peso, altura)

	definirNivel(imc)
	fmt.Println("\nSeu IMC é: ", imc)
}

func capturaNome() string {
	fmt.Println("Digite seu nome: ")

	var nome string
	fmt.Scan(&nome)

	return nome
}

func capturarPeso() float64 {
	fmt.Println("Digite seu peso: ")

	var peso float64
	fmt.Scan(&peso)

	return peso
}

func capturarAltura() float64 {
	fmt.Println("Digite sua altura: ")

	var altura float64
	fmt.Scan(&altura)

	return altura
}

func calcularImc(peso, altura float64) float64 {
	return peso / (altura * altura)
}

func definirNivel(imc float64) {
	switch {
	case imc < 18.5:
		fmt.Printf("Abaixo do peso")
	case imc >= 18.5 && imc <= 24.9:
		fmt.Printf("Peso normal")
	case imc >= 25 && imc <= 29.9:
		fmt.Printf("Sobrepeso")
	case imc >= 30 && imc <= 34.9:
		fmt.Printf("Obesidade grau I")
	case imc >= 35 && imc <= 39.9:
		fmt.Printf("Obesidade grau II")
	case imc >= 40 && imc <= 49.9:
		fmt.Printf("Obesidade grau III")
	}
}
