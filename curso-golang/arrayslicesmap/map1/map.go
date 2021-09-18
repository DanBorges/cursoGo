package main

import "fmt"

func main() {
	aprovados := make(map[int]string)

	aprovados[10110241606] = "Daniel"
	aprovados[12345678996] = "Maria"
	aprovados[78945612325] = "Ana"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[78945612325])
	delete(aprovados, 78945612325)
	fmt.Println(aprovados[78945612325])
}
