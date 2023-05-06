package model

import (
	"errors"
	"time"
)

type Compra struct {
	Mercado string
	Data    time.Time
	Itens   []ItensDaCompra
}

type ItensDaCompra struct {
	Nome string
}

func NovaCompra(mercado string, data time.Time, nomeDosItens []string) (*Compra, error) {

	if mercado == "" {
		return nil, errors.New("mercado é obrigatório")
	}

	if len(nomeDosItens) == 0 {
		return nil, errors.New("itens são obrigatórios")
	}

	var itens []ItensDaCompra

	for _, nome := range nomeDosItens {
		itens = append(itens, ItensDaCompra{Nome: nome})
	}

	return &Compra{
		Mercado: mercado,
		Data:    time.Now(),
		Itens:   itens,
	}, nil
}
