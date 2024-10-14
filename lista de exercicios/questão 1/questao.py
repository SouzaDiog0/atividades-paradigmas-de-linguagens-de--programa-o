class Carro:
    def __init__(self, marca, modelo, ano):
        self.marca = marca
        self.modelo = modelo
        self.ano = ano

    def exibir_detalhes(self):
        print(f"Marca: {self.marca}, Modelo: {self.modelo}, Ano: {self.ano}")

carro1 = Carro("MarcaA", "ModeloX", 2020)
carro2 = Carro("MarcaB", "ModeloY", 2021)
carro3 = Carro("MarcaC", "ModeloZ", 2022)

carro1.exibir_detalhes()
carro2.exibir_detalhes()
carro3.exibir_detalhes()
