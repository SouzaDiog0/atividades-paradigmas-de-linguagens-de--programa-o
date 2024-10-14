package main

import "fmt"

// Função para somar dois números
func SomarDois(a, b int) int {
    return a + b
}

// Função para somar três números
func SomarTres(a, b, c int) int {
    return a + b + c
}

func main() {
    fmt.Println(SomarDois(2, 3))    // Soma de dois números
    fmt.Println(SomarTres(2, 3, 4)) // Soma de três números
}
