package main
import "fmt"

func main() {
	var(
		n1 int 
		n2 int
		s int
	)
	fmt.Print("Informe o primeiro número: ")
	fmt.Scan(&n1)
	fmt.Print("Informe o segundo número: ")
	fmt.Scan(&n2)
	
	s = soma(n1, n2)
	fmt.Printf("A soma dos números informados é: %v\n", s)
}

func soma (x, y int) int {
		return x + y
}