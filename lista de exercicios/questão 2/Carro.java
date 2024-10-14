public class Carro {
    private String marca;
    private String modelo;
    private int ano;
    private int velocidade;

    public Carro(String marca, String modelo, int ano) {
        this.marca = marca;
        this.modelo = modelo;
        this.ano = ano;
        this.velocidade = 0;
    }

    public void acelerar() {
        velocidade += 10;
    }

    public void frear() {
        if (velocidade >= 10) {
            velocidade -= 10;
        }
    }

    public void exibirVelocidade() {
        System.out.println("Velocidade atual: " + velocidade + " km/h");
    }

    public static void main(String[] args) {
        Carro carro = new Carro("MarcaA", "ModeloX", 2020);
        carro.acelerar();
        carro.exibirVelocidade();
        carro.frear();
        carro.exibirVelocidade();
    }
}
