package main
import "fmt"

func main(){
	for{
		var a Aluno
		var qntdAlunos int
		
		for i:= range a.notas{
			fmt.Printf("Informe a %vº nota do aluno: ", i+1)
			fmt.Scanln(&a.notas[i])
		}

	}
}

type Aluno struct{
	notas [2]float64
}