package main

import "fmt"

type Carro struct {
    Marca  string
    Modelo string
    Ano    int
}

func (c Carro) ExibirDetalhes() {
    fmt.Printf("Marca: %s, Modelo: %s, Ano: %d\n", c.Marca, c.Modelo, c.Ano)
}

func main() {
    carro1 := Carro{Marca: "MarcaA", Modelo: "ModeloX", Ano: 2020}
    carro2 := Carro{Marca: "MarcaB", Modelo: "ModeloY", Ano: 2021}
    carro3 := Carro{Marca: "MarcaC", Modelo: "ModeloZ", Ano: 2022}

    carro1.ExibirDetalhes()
    carro2.ExibirDetalhes()
    carro3.ExibirDetalhes()
}
