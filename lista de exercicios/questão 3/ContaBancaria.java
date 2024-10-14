public class ContaBancaria {
    private String titular;
    private double saldo;

    public ContaBancaria(String titular, double saldoInicial) {
        this.titular = titular;
        this.saldo = saldoInicial;
    }

    public void depositar(double quantia) {
        saldo += quantia;
    }

    public void sacar(double quantia) {
        if (quantia <= saldo) {
            saldo -= quantia;
        } else {
            System.out.println("Saldo insuficiente.");
        }
    }

    public void exibirSaldo() {
        System.out.println("Saldo atual: " + saldo);
    }

    public static void main(String[] args) {
        ContaBancaria conta = new ContaBancaria("João", 500.0);
        conta.depositar(200.0);
        conta.exibirSaldo();
        conta.sacar(100.0);
        conta.exibirSaldo();
    }
}
