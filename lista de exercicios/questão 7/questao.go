package main

import "fmt"

type Professor struct {
    Nome string
}

type Escola struct {
    Nome       string
    Professores []Professor
}

func (e *Escola) AdicionarProfessor(professor Professor) {
    e.Professores = append(e.Professores, professor)
}

func (e Escola) ExibirProfessores() {
    fmt.Printf("Professores da escola %s:\n", e.Nome)
    for _, professor := range e.Professores {
        fmt.Println(professor.Nome)
    }
}

func main() {
    professor1 := Professor{Nome: "João"}
    professor2 := Professor{Nome: "Maria"}

    escola1 := Escola{Nome: "Escola A"}
    escola2 := Escola{Nome: "Escola B"}

    escola1.AdicionarProfessor(professor1)
    escola1.AdicionarProfessor(professor2)

    escola2.AdicionarProfessor(professor1)

    escola1.ExibirProfessores()
    escola2.ExibirProfessores()
}
