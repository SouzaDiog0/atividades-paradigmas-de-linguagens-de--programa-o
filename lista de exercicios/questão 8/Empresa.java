import java.util.ArrayList;
import java.util.List;

public class Empregado {
    private String nome;
    private String cargo;
    private double salario;

    public Empregado(String nome, String cargo, double salario) {
        this.nome = nome;
        this.cargo = cargo;
        this.salario = salario;
    }

    public String getNome() {
        return nome;
    }

    public String getCargo() {
        return cargo;
    }

    public double getSalario() {
        return salario;
    }
}

public class Empresa {
    private String nome;
    private List<Empregado> empregados;

    public Empresa(String nome) {
        this.nome = nome;
        this.empregados = new ArrayList<>();
    }

    public void adicionarEmpregado(Empregado empregado) {
        empregados.add(empregado);
    }

    public void exibirEmpregados() {
        System.out.println("Empregados da empresa " + nome + ":");
        for (Empregado empregado : empregados) {
            System.out.println("Nome: " + empregado.getNome() + ", Cargo: " + empregado.getCargo() + ", Salário: " + empregado.getSalario());
        }
    }

    public static void main(String[] args) {
        Empresa empresa = new Empresa("TechCorp");
        Empregado empregado1 = new Empregado("Ana", "Desenvolvedora", 5000.0);
        Empregado empregado2 = new Empregado("Carlos", "Designer", 4000.0);

        empresa.adicionarEmpregado(empregado1);
        empresa.adicionarEmpregado(empregado2);

        empresa.exibirEmpregados();
    }
}
