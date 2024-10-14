package main

import "fmt"

// Interface que define o método calcularSalario
type Funcionario interface {
    CalcularSalario() float64
}

// Struct para funcionários horistas
type FuncionarioHorista struct {
    Horas         int
    ValorPorHora  float64
}

// Implementação do método CalcularSalario para FuncionarioHorista
func (f FuncionarioHorista) CalcularSalario() float64 {
    return float64(f.Horas) * f.ValorPorHora
}

// Struct para funcionários assalariados
type FuncionarioAssalariado struct {
    SalarioFixo float64
}

// Implementação do método CalcularSalario para FuncionarioAssalariado
func (f FuncionarioAssalariado) CalcularSalario() float64 {
    return f.SalarioFixo
}

func main() {
    horista := FuncionarioHorista{Horas: 160, ValorPorHora: 50}
    assalariado := FuncionarioAssalariado{SalarioFixo: 3000}

    fmt.Printf("Salário do horista: R$%.2f\n", horista.CalcularSalario())
    fmt.Printf("Salário do assalariado: R$%.2f\n", assalariado.CalcularSalario())
}
