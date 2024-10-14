public class Animal {
    public void som() {
        System.out.println("Som de animal");
    }
}

public class Cachorro extends Animal {
    @Override
    public void som() {
        System.out.println("O cachorro faz: Au Au");
    }
}

public class Gato extends Animal {
    @Override
    public void som() {
        System.out.println("O gato faz: Miau");
    }
}

public class Main {
    public static void emitirSons(Animal[] animais) {
        for (Animal animal : animais) {
            animal.som();
        }
    }

    public static void main(String[] args) {
        Animal[] animais = { new Cachorro(), new Gato() };
        emitirSons(animais);
    }
}
