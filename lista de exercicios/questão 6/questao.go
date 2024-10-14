package main

import "fmt"

type Motor struct {
    Potencia int
}

type Carro struct {
    Marca  string
    Modelo string
    Ano    int
    Motor  Motor
}

func (c Carro) ExibirDetalhes() {
    fmt.Printf("Marca: %s, Modelo: %s, Ano: %d, Potência do motor: %d cavalos\n", c.Marca, c.Modelo, c.Ano, c.Motor.Potencia)
}

func main() {
    motor := Motor{Potencia: 200}
    carro := Carro{Marca: "MarcaA", Modelo: "ModeloX", Ano: 2020, Motor: motor}
    carro.ExibirDetalhes()
}
