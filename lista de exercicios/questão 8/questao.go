package main

import "fmt"

type Empregado struct {
    Nome   string
    Cargo  string
    Salario float64
}

type Empresa struct {
    Nome      string
    Empregados []Empregado
}

func (e *Empresa) AdicionarEmpregado(empregado Empregado) {
    e.Empregados = append(e.Empregados, empregado)
}

func (e Empresa) ExibirEmpregados() {
    fmt.Printf("Empregados da empresa %s:\n", e.Nome)
    for _, empregado := range e.Empregados {
        fmt.Printf("Nome: %s, Cargo: %s, Salário: %.2f\n", empregado.Nome, empregado.Cargo, empregado.Salario)
    }
}

func main() {
    empresa := Empresa{Nome: "TechCorp"}
    empregado1 := Empregado{Nome: "Ana", Cargo: "Desenvolvedora", Salario: 5000.0}
    empregado2 := Empregado{Nome: "Carlos", Cargo: "Designer", Salario: 4000.0}

    empresa.AdicionarEmpregado(empregado1)
    empresa.AdicionarEmpregado(empregado2)

    empresa.ExibirEmpregados()
}
