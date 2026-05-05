package main
import "fmt"

func main(){
	vetorOriginal:=make([]int, 30)
	vetorModificado := make([]int, 30)

	for i := range vetorOriginal {
		fmt.Print("Informe um valor inteiro: ")
		fmt.Scanln(&vetorOriginal[i])

		if i%2 == 0 {
			vetorModificado[i] = vetorOriginal[i] * 2
		} else {
			vetorModificado[i] = vetorOriginal[i] * 3
		}
	}
	fmt.Printf("Vetor original: %v\n", vetorOriginal)
	fmt.Printf("Vetor modificado: %v\n", vetorModificado) 
}