package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d) \n", pessoa, texto, qtde)
	}
}

func main() {
	//fale("Maria", "pq vc nao fala comigo", 3)
	//fale("João", "Só posso falar depois de vc", 1)

	go fale("Maria", "pq vc nao fala comigo", 500)
	go fale("João", "Só posso falar depois de vc", 500)
	//fmt.Println("FIM")
	time.Sleep(time.Second * 5)
}
