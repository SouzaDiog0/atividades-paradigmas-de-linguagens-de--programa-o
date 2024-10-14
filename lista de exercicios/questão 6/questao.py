class Motor:
    def __init__(self, potencia):
        self.potencia = potencia

class Carro:
    def __init__(self, marca, modelo, ano, motor):
        self.marca = marca
        self.modelo = modelo
        self.ano = ano
        self.motor = motor

    def exibir_detalhes(self):
        print(f"Marca: {self.marca}, Modelo: {self.modelo}, Ano: {self.ano}, Potência do motor: {self.motor.potencia} cavalos")

motor = Motor(200)
carro = Carro("MarcaA", "ModeloX", 2020, motor)
carro.exibir_detalhes()
