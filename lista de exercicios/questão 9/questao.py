from abc import ABC, abstractmethod

class Imprimivel(ABC):
    @abstractmethod
    def imprimir(self):
        pass

class Relatorio(Imprimivel):
    def __init__(self, conteudo):
        self.conteudo = conteudo

    def imprimir(self):
        print(f"Imprimindo Relatório: {self.conteudo}")

class Contrato(Imprimivel):
    def __init__(self, conteudo):
        self.conteudo = conteudo

    def imprimir(self):
        print(f"Imprimindo Contrato: {self.conteudo}")

relatorio = Relatorio("Relatório de Vendas")
contrato = Contrato("Contrato de Trabalho")

relatorio.imprimir()
contrato.imprimir()
