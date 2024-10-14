abstract class Funcionario {
    public abstract double calcularSalario();
}

class FuncionarioHorista extends Funcionario {
    private int horas;
    private double valorPorHora;

    public FuncionarioHorista(int horas, double valorPorHora) {
        this.horas = horas;
        this.valorPorHora = valorPorHora;
    }

    @Override
    public double calcularSalario() {
        return horas * valorPorHora;
    }
}

class FuncionarioAssalariado extends Funcionario {
    private double salarioFixo;

    public FuncionarioAssalariado(double salarioFixo) {
        this.salarioFixo = salarioFixo;
    }

    @Override
    public double calcularSalario() {
        return salarioFixo;
    }
}

public class Main {
    public static void main(String[] args) {
        Funcionario horista = new FuncionarioHorista(160, 50);
        Funcionario assalariado = new FuncionarioAssalariado(3000);

        System.out.println("Salário do horista: R$" + horista.calcularSalario());
        System.out.println("Salário do assalariado: R$" + assalariado.calcularSalario());
    }
}
