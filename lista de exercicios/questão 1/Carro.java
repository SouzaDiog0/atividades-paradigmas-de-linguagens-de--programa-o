public class Carro {
    private String marca;
    private String modelo;
    private int ano;

    public Carro(String marca, String modelo, int ano) {
        this.marca = marca;
        this.modelo = modelo;
        this.ano = ano;
    }

    public void exibirDetalhes() {
        System.out.println("Marca: " + marca + ", Modelo: " + modelo + ", Ano: " + ano);
    }

    public static void main(String[] args) {
        Carro carro1 = new Carro("MarcaA", "ModeloX", 2020);
        Carro carro2 = new Carro("MarcaB", "ModeloY", 2021);
        Carro carro3 = new Carro("MarcaC", "ModeloZ", 2022);
        
        carro1.exibirDetalhes();
        carro2.exibirDetalhes();
        carro3.exibirDetalhes();
    }
}
