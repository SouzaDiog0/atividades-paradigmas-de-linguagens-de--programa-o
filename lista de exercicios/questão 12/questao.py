class Produto:
    def __init__(self, nome, preco):
        self.nome = nome
        self.preco = preco

    # Sobrecarga do operador +
    def __add__(self, outro):
        return Produto(f"{self.nome} & {outro.nome}", self.preco + outro.preco)

    def __str__(self):
        return f"Produto: {self.nome}, Preço: R${self.preco:.2f}"

# Exemplo de uso
produto1 = Produto("Camiseta", 50.00)
produto2 = Produto("Boné", 30.00)
produto3 = produto1 + produto2

print(produto1)
print(produto2)
print(produto3)  # Produto combinado com preço somado
