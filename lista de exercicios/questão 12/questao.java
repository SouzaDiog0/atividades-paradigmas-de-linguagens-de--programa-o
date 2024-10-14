class Produto {
    private String nome;
    private double preco;

    public Produto(String nome, double preco) {
        this.nome = nome;
        this.preco = preco;
    }

    public Produto somar(Produto outro) {
        return new Produto(this.nome + " & " + outro.nome, this.preco + outro.preco);
    }

    @Override
    public String toString() {
        return String.format("Produto: %s, Preço: R$%.2f", nome, preco);
    }

    public static void main(String[] args) {
        Produto produto1 = new Produto("Camiseta", 50.00);
        Produto produto2 = new Produto("Boné", 30.00);
        Produto produto3 = produto1.somar(produto2);

        System.out.println(produto1);
        System.out.println(produto2);
        System.out.println(produto3);
    }
}
