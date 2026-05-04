package main
import "fmt"

func main(){
	vetorInt := make([]int, 10)
	x, p := 0, 0

	for i := range vetorInt{
		fmt.Print("Informe um valor inteiro: ")
		fmt.Scanln(&vetorInt[i])
		
		if i == 0 {
			x = vetorInt[i]
			p = 0
		} else if x > vetorInt[i]{
			x = vetorInt[i]
			p = i
		}
	}
	fmt.Printf("O menor elemento do vetor é %v, e sua posição dentro do vetor é %v\n", x, p)
}
