import java.util.ArrayList;
import java.util.List;

public class Professor {
    private String nome;

    public Professor(String nome) {
        this.nome = nome;
    }

    public String getNome() {
        return nome;
    }
}

public class Escola {
    private String nome;
    private List<Professor> professores;

    public Escola(String nome) {
        this.nome = nome;
        this.professores = new ArrayList<>();
    }

    public void adicionarProfessor(Professor professor) {
        professores.add(professor);
    }

    public void exibirProfessores() {
        System.out.println("Professores da escola " + nome + ":");
        for (Professor professor : professores) {
            System.out.println(professor.getNome());
        }
    }

    public static void main(String[] args) {
        Professor professor1 = new Professor("João");
        Professor professor2 = new Professor("Maria");

        Escola escola1 = new Escola("Escola A");
        Escola escola2 = new Escola("Escola B");

        escola1.adicionarProfessor(professor1);
        escola1.adicionarProfessor(professor2);

        escola2.adicionarProfessor(professor1);

        escola1.exibirProfessores();
        escola2.exibirProfessores();
    }
}
