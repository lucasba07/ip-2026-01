package main
import "fmt"

func main() {
	var (
		num [5]int
		soma int = 0
	)
	
	for i := range num{
		fmt.Printf("Informe um valor qualquer: ")
		fmt.Scan(&num[i])
		soma += num[i]
	}
	fmt.Printf("A soma dos valores informados é: %v\n", soma)
}