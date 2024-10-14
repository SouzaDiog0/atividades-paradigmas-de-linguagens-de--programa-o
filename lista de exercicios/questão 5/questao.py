class Animal:
    def som(self):
        print("Som de animal")

class Cachorro(Animal):
    def som(self):
        print("O cachorro faz: Au Au")

class Gato(Animal):
    def som(self):
        print("O gato faz: Miau")

def emitir_sons(animais):
    for animal in animais:
        animal.som()

animais = [Cachorro(), Gato()]
emitir_sons(animais)
