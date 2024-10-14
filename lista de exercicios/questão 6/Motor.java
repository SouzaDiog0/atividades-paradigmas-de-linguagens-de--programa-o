public class Motor {
    private int potencia;

    public Motor(int potencia) {
        this.potencia = potencia;
    }

    public int getPotencia() {
        return potencia;
    }
}

public class Carro {
    private String marca;
    private String modelo;
    private int ano;
    private Motor motor;

    public Carro(String marca, String modelo, int ano, Motor motor) {
        this.marca = marca;
        this.modelo = modelo;
        this.ano = ano;
        this.motor = motor;
    }

    public void exibirDetalhes() {
        System.out.println("Marca: " + marca + ", Modelo: " + modelo + ", Ano: " + ano + ", Potência do motor: " + motor.getPotencia() + " cavalos");
    }

    public static void main(String[] args) {
        Motor motor = new Motor(200);
        Carro carro = new Carro("MarcaA", "ModeloX", 2020, motor);
        carro.exibirDetalhes();
    }
}
