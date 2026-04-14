package main
import "fmt"

func main() {
	var (
		num [3]int
		maior int
	)	
	maior = maiorValor(num, maior)
	fmt.Printf("O maior valor informado é: %v\n", maior)
}

func maiorValor (x [3]int, y int) int {
	for i := range x{
		fmt.Printf("Informe o %vº número: ", i+1)
		fmt.Scan(&x[i])
		if (i == 0){
			y = x[i]
		} else{
			if x[i] > y{
				y = x[i]
			}
		}
	}
	return y
}