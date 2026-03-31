package main
import "fmt"

func main() {
	var(
		nota [3]float64
		soma float64 = 0
	)
	
	for i := range nota{
		fmt.Printf("Digite a %vº nota do aluno: ", (i+1))
		fmt.Scan(&nota[i])
		soma += nota[i]
	}
	
	media := soma/ 3
	fmt.Printf("A média geral do aluno é: %.2f\n", media)

	for i := range nota{
		if nota[i] > media{
			fmt.Printf("A %vº nota, %.2f, está acima da média geral do aluno, que é %.2f\n", (i+1), nota[i], media)
		}
	}       
}