package main
import "fmt"

func main(){
	vetor1:= make([]int, 10)
	vetor2:= make([]int, 5)
	vetorResultante1:= make([]int, 0, 10)
	vetorResultante2:= make([]int, 0, 10)
	somaVetor2:=0

	for i := range vetor1 {
		fmt.Print("Digite um valor para o vetor1: ")
		fmt.Scanln(&vetor1[i])
	}

	for i := range vetor2 {
		fmt.Print("Digite um valor para o vetor2: ")
		fmt.Scanln(&vetor2[i])
		somaVetor2 += vetor2[i]
	}

	for i:= range vetor1{
		if vetor1[i]%2==0{
			vetorResultante1 = append(vetorResultante1, (vetor1[i] + somaVetor2))
		} else{
			vetorResultante2 = append(vetorResultante2, (vetor1[i] + somaVetor2))
		}
	}
	fmt.Printf("\nPrimeiro vetor: %v\n", vetor1)
	fmt.Printf("Segundo vetor: %v\n", vetor2)
	fmt.Printf("Primeiro vetor resultante: %v\n", vetorResultante1)
	fmt.Printf("Segundo vetor resultante: %v\n", vetorResultante2)
}
