package main
import (
	"fmt"
	"sort"
)

func main(){
	vetor := make([]int, 10)

	for i:=range vetor{
		fmt.Print("Informe um valor inteiro: ")
		fmt.Scanln(&vetor[i])
	}
	fmt.Printf("Vetor original: %v\n", vetor)
	sort.Ints(vetor)
	fmt.Printf("Vetor em ordem crescente: %v\n", vetor)
}