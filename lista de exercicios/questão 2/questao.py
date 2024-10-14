class Carro:
    def __init__(self, marca, modelo, ano):
        self.marca = marca
        self.modelo = modelo
        self.ano = ano
        self.velocidade = 0

    def acelerar(self):
        self.velocidade += 10

    def frear(self):
        if self.velocidade >= 10:
            self.velocidade -= 10

    def exibir_velocidade(self):
        print(f"Velocidade atual: {self.velocidade} km/h")

carro = Carro("MarcaA", "ModeloX", 2020)
carro.acelerar()
carro.exibir_velocidade()
carro.frear()
carro.exibir_velocidade()
