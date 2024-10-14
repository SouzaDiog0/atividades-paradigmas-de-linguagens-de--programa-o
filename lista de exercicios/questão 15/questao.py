class SaldoInsuficienteException(Exception):
    def __init__(self, saldo, valor):
        super().__init__(f"Saldo insuficiente. Saldo atual: R${saldo}, Tentativa de saque: R${valor}")

class ContaBancaria:
    def __init__(self, titular, saldo=0):
        self.titular = titular
        self.saldo = saldo

    def sacar(self, valor):
        if valor > self.saldo:
            raise SaldoInsuficienteException(self.saldo, valor)
        self.saldo -= valor
        print(f"Saque de R${valor} realizado com sucesso.")

# Exemplo de uso
conta = ContaBancaria("João", 100)

try:
    conta.sacar(200)  # Exceção será lançada aqui
except SaldoInsuficienteException as e:
    print(e)
