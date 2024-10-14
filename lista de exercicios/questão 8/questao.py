class Empregado:
    def __init__(self, nome, cargo, salario):
        self.nome = nome
        self.cargo = cargo
        self.salario = salario

class Empresa:
    def __init__(self, nome):
        self.nome = nome
        self.empregados = []

    def adicionar_empregado(self, empregado):
        self.empregados.append(empregado)

    def exibir_empregados(self):
        print(f"Empregados da empresa {self.nome}:")
        for empregado in self.empregados:
            print(f"Nome: {empregado.nome}, Cargo: {empregado.cargo}, Salário: {empregado.salario}")

empresa = Empresa("TechCorp")
empregado1 = Empregado("Ana", "Desenvolvedora", 5000.0)
empregado2 = Empregado("Carlos", "Designer", 4000.0)

empresa.adicionar_empregado(empregado1)
empresa.adicionar_empregado(empregado2)

empresa.exibir_empregados()
