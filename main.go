package main

import (
	"fmt"
	"goexer/model"
	"time"
)

func main() {
	var nomeDosItens []string
	nomeDosItens = append(nomeDosItens, "arroz")
	nomeDosItens = append(nomeDosItens, "feijao")
	nomeDosItens = append(nomeDosItens, "carne")
	nomeDosItens = append(nomeDosItens, "sabonete")

	compra, err := model.NovaCompra("Mercadinho", time.Now(), nomeDosItens)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(compra)
	}
}
