from abc import ABC, abstractmethod

class Funcionario(ABC):
    @abstractmethod
    def calcular_salario(self):
        pass

class FuncionarioHorista(Funcionario):
    def __init__(self, horas, valor_por_hora):
        self.horas = horas
        self.valor_por_hora = valor_por_hora

    def calcular_salario(self):
        return self.horas * self.valor_por_hora

class FuncionarioAssalariado(Funcionario):
    def __init__(self, salario_fixo):
        self.salario_fixo = salario_fixo

    def calcular_salario(self):
        return self.salario_fixo

# Exemplo de uso
horista = FuncionarioHorista(160, 50)
assalariado = FuncionarioAssalariado(3000)

print(f"Salário do horista: R${horista.calcular_salario()}")
print(f"Salário do assalariado: R${assalariado.calcular_salario()}")
