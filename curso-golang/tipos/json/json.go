package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {

	// struct para Json
	p1 := produto{1, "Notebook", 5879.87, []string{"Promoção", "Eletrônco"}}
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))

	// Json para Struct
	var p2 produto
	jsonToString := `{"id":2, "nome":"Caneta", "preco":8.90, "tags":["papelaria", "importado"] }`
	json.Unmarshal([]byte(jsonToString), &p2)
	fmt.Println(p2.Tags[1])
	fmt.Println(p2.Nome)
	fmt.Println(p2.Preco)

}
