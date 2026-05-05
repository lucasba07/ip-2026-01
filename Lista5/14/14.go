package main
import "fmt"

func main(){
	vetor1 := make([]int, 10)
	vetor2 := make([]int, 10)
	vetorR := make([]int, 0, 20)

	for i:= range vetor1{
		fmt.Print("Informe um valor inteiro para o vetor 1: ")
		fmt.Scanln(&vetor1[i])

		fmt.Print("Informe um valor inteiro para o vetor 2: ")
		fmt.Scanln(&vetor2[i])

		vetorR = append(vetorR, vetor1[i], vetor2[i])
	}
	fmt.Printf("Vetor 1: %v\n", vetor1)
	fmt.Printf("Vetor 2: %v\n", vetor2)
	fmt.Printf("Vetor resultante da intercalação: %v\n", vetorR)
}
