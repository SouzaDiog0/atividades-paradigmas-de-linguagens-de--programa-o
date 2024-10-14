package main

import "fmt"

type Carro struct {
    Marca     string
    Modelo    string
    Ano       int
    Velocidade int
}

func (c *Carro) Acelerar() {
    c.Velocidade += 10
}

func (c *Carro) Frear() {
    if c.Velocidade >= 10 {
        c.Velocidade -= 10
    }
}

func (c Carro) ExibirVelocidade() {
    fmt.Printf("Velocidade atual: %d km/h\n", c.Velocidade)
}

func main() {
    carro := Carro{Marca: "MarcaA", Modelo: "ModeloX", Ano: 2020, Velocidade: 0}
    carro.Acelerar()
    carro.ExibirVelocidade()
    carro.Frear()
    carro.ExibirVelocidade()
}
