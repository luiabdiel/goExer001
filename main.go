package main

import (
	"fmt"
	"time"
)

func main() {
	var itens []itensDaCompra

	itens = append(itens, itensDaCompra{nome: "arroz"})
	itens = append(itens, itensDaCompra{nome: "feijao"})

	compra := compra{
		mercado: "mercado x",
		data:    time.Now(),
		itens:   itens,
	}

	fmt.Println(compra)
}
