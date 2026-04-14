package main
import "fmt"

func main() {
	var(
		num [3]float64
		soma float64
	)
	fmt.Printf("A média dos valores informados é: %.2f\n", media(num, soma))
}

func media (x [3]float64, y float64) float64 {
	for i := range x {
		fmt.Printf("Informe o %vº valor qualquer: ", i+1)
		fmt.Scan(&x[i])
		y += x[i]
	}
	return y / 3
}