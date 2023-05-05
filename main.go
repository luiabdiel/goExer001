package main

import (
	"fmt"
	"goexer/model"
	"time"
)

func main() {
	var itens []model.ItensDaCompra

	itens = append(itens, model.ItensDaCompra{Nome: "arroz"})
	itens = append(itens, model.ItensDaCompra{Nome: "feijao"})
	itens = append(itens, model.ItensDaCompra{Nome: "carne"})

	compra := model.Compra{
		Mercado: "mercado x",
		Data:    time.Now(),
		Itens:   itens,
	}

	fmt.Println(compra)
}
