package main

import "fmt"

func main() {
	i := 1

	var p *int = nil
	p = &i //pegando o endereço da variavel
	*p++   //desrreferenciando (pegando o valor)
	i++

	fmt.Println(p, *p, i, &i)
}
