package main

import "fmt"

func main() {
	funcESalarios := map[string]float64{
		"José Maria": 1000.00,
		"Pedro Hugo": 2500,
		"Zé":         14000.98,
	}

	funcESalarios["Daniel Lemes"] = 5890.55

	delete(funcESalarios, "inexistente")

	for nome, salario := range funcESalarios {
		fmt.Println(nome, salario)
	}
}
