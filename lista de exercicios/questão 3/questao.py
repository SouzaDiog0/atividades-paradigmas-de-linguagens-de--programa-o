class ContaBancaria:
    def __init__(self, titular, saldo_inicial):
        self.__titular = titular
        self.__saldo = saldo_inicial

    def depositar(self, quantia):
        self.__saldo += quantia

    def sacar(self, quantia):
        if quantia <= self.__saldo:
            self.__saldo -= quantia
        else:
            print("Saldo insuficiente.")

    def exibir_saldo(self):
        print(f"Saldo atual: {self.__saldo}")

conta = ContaBancaria("João", 500.0)
conta.depositar(200.0)
conta.exibir_saldo()
conta.sacar(100.0)
conta.exibir_saldo()
