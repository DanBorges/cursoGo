package main

import "fmt"

func multiplicacao(a, b int) int {
	return a * b
}

func exex(funcao func(int, int) int, p1, p2 int) int {
	return funcao(p1, p2)
}

func main() {
	resultado := exex(multiplicacao, 3, 4)
	fmt.Println(resultado)
}
