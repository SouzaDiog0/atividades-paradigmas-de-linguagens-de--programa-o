class Animal:
    def som(self):
        print("Som de animal")

class Cachorro(Animal):
    def som(self):
        print("O cachorro faz: Au Au")

class Gato(Animal):
    def som(self):
        print("O gato faz: Miau")

cachorro = Cachorro()
gato = Gato()

cachorro.som()
gato.som()
