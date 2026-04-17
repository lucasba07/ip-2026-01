package main
import "fmt"

func main(){
    var(
        qntdAlunos int
        alunos []Aluno
        qntdR int
        qntdE int
        qntdA int
        somaMedia float64
    )

    fmt.Print("A sala contém quantos alunos? ")
    fmt.Scanln(&qntdAlunos)
    alunos = make([]Aluno, qntdAlunos)

    for i:= range alunos{
        var a Aluno
        fmt.Printf("Informe a primeira nota do %vº aluno: ", (i+1))
        fmt.Scanln(&a.n1)

        fmt.Printf("Informe a segunda nota do %vº aluno: ", (i+1))
        fmt.Scanln(&a.n2)

        a.media = (a.n1 + a.n2) / 2
        somaMedia += a.media
        alunos[i] = a
        res:= resultadoMedia(a.media)

        fmt.Printf("Resultado do %vº aluno: %v\n", (i+1), res)

        switch res {
            case "REPROVADO":
                qntdR++
            case "EXAME":
                qntdE++
            default:
                qntdA++
        }
    }
    mediaGeral:= somaMedia / float64(qntdAlunos)

    fmt.Printf("MÉDIA DA TURMA: %.2f | TOTAL DE APROVADOS: %v | TOTAL DE REPROVADOS: %v | TOTAL DE ALUNOS DE EXAME: %v\n", mediaGeral, qntdA, qntdR, qntdE)
}

type Aluno struct{
    n1 float64
    n2 float64
    media float64
}

func resultadoMedia (media float64) string{
    switch{
        case media <= 3:
            return "REPROVADO"
        case media > 3 && media < 7:
            return "EXAME"
        default:
            return "APROVADO"
    }
}