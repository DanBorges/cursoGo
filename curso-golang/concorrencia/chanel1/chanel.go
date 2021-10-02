package main

import "fmt"

func main() {

	ch := make(chan int, 1)
	ch <- 1 //Enviando dados para um canal
	<-ch    //Recebendo dados de um canal

	ch <- 2 //Enviando dados para um canal
	fmt.Println(<-ch)
}
