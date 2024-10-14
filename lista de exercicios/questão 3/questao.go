package main

import "fmt"

type ContaBancaria struct {
    Titular string
    saldo   float64
}

func (c *ContaBancaria) Depositar(quantia float64) {
    c.saldo += quantia
}

func (c *ContaBancaria) Sacar(quantia float64) {
    if quantia <= c.saldo {
        c.saldo -= quantia
    } else {
        fmt.Println("Saldo insuficiente.")
    }
}

func (c ContaBancaria) ExibirSaldo() {
    fmt.Printf("Saldo atual: %.2f\n", c.saldo)
}

func main() {
    conta := ContaBancaria{Titular: "João", saldo: 500.0}
    conta.Depositar(200.0)
    conta.ExibirSaldo()
    conta.Sacar(100.0)
    conta.ExibirSaldo()
}
