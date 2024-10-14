package main

import (
    "fmt"
    "sync"
)

// Struct para configuração
type Configuracao struct{}

var instancia *Configuracao
var once sync.Once

// Função para obter a instância única
func GetInstancia() *Configuracao {
    once.Do(func() {
        instancia = &Configuracao{}
    })
    return instancia
}

func main() {
    config1 := GetInstancia()
    config2 := GetInstancia()

    fmt.Println(config1 == config2)  // Saída: true (mesma instância)
}
