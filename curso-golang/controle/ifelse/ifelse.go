package main

import "fmt"

func imprimirResultado(nota float64) {
	if nota >= 7 {
		fmt.Println("O Aluno foi aprovado com a nota", nota)
	} else {
		fmt.Println("O Aluno foi reprovado com a nora", nota)
	}
}

func main() {
	imprimirResultado(7)
	imprimirResultado(5.9)
}
