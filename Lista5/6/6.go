package main
import "fmt"

func main(){
	vetor := make([]int, 100)

	for i := 99; i >= 0; i--{
		vetor[i] = (i+1)
		fmt.Printf("%v ", vetor[i])
	}
}
