class Calculadora:
    def somar(self, a, b, c=0):
        return a + b + c

# Exemplo de uso:
calc = Calculadora()
print(calc.somar(2, 3))        # Soma de dois números
print(calc.somar(2, 3, 4))     # Soma de três números
