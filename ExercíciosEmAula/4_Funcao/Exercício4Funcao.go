package main
import "fmt"

func main() {
	var num float64
	fmt.Print("Informe um número qualquer: ")
	fmt.Scan(&num)
	fmt.Printf("o fatorial de %v é: %.2f", num, fatorial(num))
}

func fatorial (x float64) float64 {
	if x == 0 || x == 1 {
		return 1
	} else {
		return (x * fatorial(x-1))
	}
}
