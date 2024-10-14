public interface Imprimivel {
    void imprimir();
}

public class Relatorio implements Imprimivel {
    private String conteudo;

    public Relatorio(String conteudo) {
        this.conteudo = conteudo;
    }

    @Override
    public void imprimir() {
        System.out.println("Imprimindo Relatório: " + conteudo);
    }
}

public class Contrato implements Imprimivel {
    private String conteudo;

    public Contrato(String conteudo) {
        this.conteudo = conteudo;
    }

    @Override
    public void imprimir() {
        System.out.println("Imprimindo Contrato: " + conteudo);
    }
}

public class Main {
    public static void main(String[] args) {
        Imprimivel relatorio = new Relatorio("Relatório de Vendas");
        Imprimivel contrato = new Contrato("Contrato de Trabalho");

        relatorio.imprimir();
        contrato.imprimir();
    }
}
