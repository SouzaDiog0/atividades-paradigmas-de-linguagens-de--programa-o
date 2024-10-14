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

func main() {
    var animal Animal

    animal = Cachorro{}
    animal.Som()

    animal = Gato{}
    animal.Som()
}
