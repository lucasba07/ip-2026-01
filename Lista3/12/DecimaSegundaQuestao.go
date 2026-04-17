package main
import "fmt"

func main(){
	var(
		num float64
		soma float64
	)
	
	fmt.Print("Digite o valor de X: ")
	fmt.Scanln(&num)

	lista := make ([]float64, 20)

	for i:= range lista{
		switch{
			case i == 0:
				lista[i] = num
			case i % 2 != 0:
				lista[i] = -(num/float64(fatorial(i)))
			default:
				lista[i] = (num/float64(fatorial(i)))
		}
		soma += lista[i]
	}
	fmt.Printf("A soma resulta em: %.2f\n", soma)
}

func fatorial (num int) int {
	if num == 0 {
		return 1
	} 
	return num * fatorial(num-1)
}
