package main
import "fmt"

func main(){
	a:= make([]int, 10)

	for i:=range a{
		fmt.Print("Informe um número inteiro: ")
		fmt.Scanln(&a[i])
	}
	verificaRepetido(a)
}

func verificaRepetido(a []int){
	countMap := make(map[int]int)
	for _, v := range a {
		countMap[v]++
	}
	vetorRepetidos := make([]int, 0)
	contRepetidos := make([]int, 0)
	for k, v := range countMap {
		if v > 1 {
			vetorRepetidos = append(vetorRepetidos, k)
			contRepetidos = append(contRepetidos, v)
		}
	}
	for i := range vetorRepetidos {
		fmt.Printf("Número %d aparece %d vezes\n", vetorRepetidos[i], contRepetidos[i])
	}
}
