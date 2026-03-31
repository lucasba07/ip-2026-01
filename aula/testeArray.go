package main
import "fmt"

func main() {
	const numNotas int = 3 
	var (
		nota [numNotas]float64
		soma float64 = 0
	)
	
	for i := range numNotas {
		fmt.Printf("Digite a nota %d: ", i+1)
		fmt.Scan(&nota[i])
		soma += nota[i]
	}

	for i := range nota {
		fmt.Printf("Nota %d: %.2f\n", i+1, nota[i])
	}

	fmt.Printf("Soma das notas: %f\n", soma)
}