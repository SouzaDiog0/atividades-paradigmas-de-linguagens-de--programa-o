class SaldoInsuficienteException extends Exception {
    public SaldoInsuficienteException(double saldo, double valor) {
        super(String.format("Saldo insuficiente. Saldo atual: R$%.2f, Tentativa de saque: R$%.2f", saldo, valor));
    }
}

class ContaBancaria {
    private String titular;
    private double saldo;

    public ContaBancaria(String titular, double saldo) {
        this.titular = titular;
        this.saldo = saldo;
    }

    public void sacar(double valor) throws SaldoInsuficienteException {
        if (valor > saldo) {
            throw new SaldoInsuficienteException(saldo, valor);
        }
        saldo -= valor;
        System.out.printf("Saque de R$%.2f realizado com sucesso.%n", valor);
    }
    
    public static void main(String[] args) {
        ContaBancaria conta = new ContaBancaria("João", 100);

        try {
            conta.sacar(200);  // Exceção será lançada aqui
        } catch (SaldoInsuficienteException e) {
            System.out.println(e.getMessage());
        }
    }
}
