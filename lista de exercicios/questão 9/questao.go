package main

import "fmt"

type Imprimivel interface {
    Imprimir()
}

type Relatorio struct {
    Conteudo string
}

func (r Relatorio) Imprimir() {
    fmt.Printf("Imprimindo Relatório: %s\n", r.Conteudo)
}

type Contrato struct {
    Conteudo string
}

func (c Contrato) Imprimir() {
    fmt.Printf("Imprimindo Contrato: %s\n", c.Conteudo)
}

func main() {
    var documento Imprimivel

    documento = Relatorio{Conteudo: "Relatório de Vendas"}
    documento.Imprimir()

    documento = Contrato{Conteudo: "Contrato de Trabalho"}
    documento.Imprimir()
}
