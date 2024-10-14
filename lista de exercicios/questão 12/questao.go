package main

import (
    "fmt"
)

// Struct Produto
type Produto struct {
    Nome  string
    Preco float64
}

// Função que combina dois produtos
func Somar(prod1, prod2 Produto) Produto {
    nomeCombinado := prod1.Nome + " & " + prod2.Nome
    precoCombinado := prod1.Preco + prod2.Preco
    return Produto{Nome: nomeCombinado, Preco: precoCombinado}
}

func main() {
    produto1 := Produto{Nome: "Camiseta", Preco: 50.00}
    produto2 := Produto{Nome: "Boné", Preco: 30.00}
    produto3 := Somar(produto1, produto2)

    fmt.Printf("Produto: %s, Preço: R$%.2f\n", produto1.Nome, produto1.Preco)
    fmt.Printf("Produto: %s, Preço: R$%.2f\n", produto2.Nome, produto2.Preco)
    fmt.Printf("Produto: %s, Preço: R$%.2f\n", produto3.Nome, produto3.Preco)
}
