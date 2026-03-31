package main
import ("fmt")

func main(){
	var nota, media, soma float64
	var qntd int

	fmt.Print("A classe tem quantos alunos? \n")
	fmt.Scan(&qntd)

	for x:= 1; x <= qntd; x++{
		fmt.Printf("Informe a nota do %vº aluno: \n", x)
		fmt.Scan(&nota)
		soma += nota
	}
	media = soma / float64(qntd)
	fmt.Printf("A média da turma é: %.2f\n", media)
}