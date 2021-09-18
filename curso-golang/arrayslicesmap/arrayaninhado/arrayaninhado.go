package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 15785.75,
			"Guga Pereira":   12000,
		},
		"J": {
			"Julio Borges": 5963,
		},
		"P": {
			"Paulo Henrique": 10000.85,
		},
	}

	delete(funcsPorLetra, "P")
	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}
