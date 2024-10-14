package main

import "fmt"

type Animal interface {
    Som()
}

type Cachorro struct{}

func (Cachorro) Som() {
    fmt.Println("O cachorro faz: Au Au")
}

type Gato struct{}

func (Gato) Som() {
    fmt.Println("O gato faz: Miau")
}

func EmitirSons(animais []Animal) {
    for _, animal := range animais {
        animal.Som()
    }
}

func main() {
    animais := []Animal{Cachorro{}, Gato{}}
    EmitirSons(animais)
}
