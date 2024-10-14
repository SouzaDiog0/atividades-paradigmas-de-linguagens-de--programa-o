package main

import (
    "errors"
    "fmt"
)

// Função que cria um erro personalizado
func SaldoInsuficienteError(saldo, valor float64) error {
    return errors.New(fmt.Sprintf("Saldo insuficiente. Saldo atual: R$%.2f, Tentativa de saque: R$%.2f", saldo, valor))
}

type ContaBancaria struct {
    Titular string
    Saldo   float64
}

// Método para realizar saque
func (c *ContaBancaria) Sacar(valor float64) error {
    if valor > c.Saldo {
        return SaldoInsuficienteError(c.Saldo, valor)
    }
    c.Saldo -= valor
    fmt.Printf("Saque de R$%.2f realizado com sucesso.\n", valor)
    return nil
}

func main() {
    conta := ContaBancaria{Titular: "João", Saldo: 100}

    if err := conta.Sacar(200); err != nil {
        fmt.Println(err)  // Exibe o erro personalizado
    }
}
