package main
import "fmt"

func main(){
	vetor := make([]int, 100)
	cont := 1

	for i := range vetor{
		vetor[i] = cont
		cont += 2
		fmt.Printf("%v ", vetor[i])
	}
}
