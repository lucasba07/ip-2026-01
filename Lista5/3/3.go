package main
import "fmt"

func main(){
	vetor := make([]int, 10)
	pares := make([]int, 0, 10)
	impares := make([]int, 0, 10)
	somaPar, qntdImpar := 0, 0

	for i:=range vetor{
		fmt.Print("Informe um valor inteiro: ")
		fmt.Scanln(&vetor[i])

		if vetor[i]%2==0{
			pares = append(pares, vetor[i])
			somaPar += vetor[i]
		} else{
			impares = append(impares, vetor[i])
			qntdImpar += 1
		}
	}
	fmt.Printf("\nPares digitados: %v\n", pares)
	fmt.Printf("Soma Par: %v\n", somaPar)
	fmt.Printf("Ímpares digitados: %v\n", impares)
	fmt.Printf("Quantidade de ímpares digitados: %v\n", qntdImpar)
}


