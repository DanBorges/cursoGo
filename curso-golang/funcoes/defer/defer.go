package main

import "fmt"

func obterValorAprovado(numero int) int {
	defer fmt.Println("FIM")

	if numero > 5000 {
		fmt.Println("Número Alto")
		return 5000
	} else {
		fmt.Println("Número Baixo")
		return numero
	}
}

func main() {
	fmt.Println(obterValorAprovado(6000))
	fmt.Println(obterValorAprovado(3000))
}
