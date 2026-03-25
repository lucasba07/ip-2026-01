package main
import ("fmt")

func main(){
	var a, b, c float32

	fmt.Print("Informe o valor do coeficiente A: ")
	fmt.Scan(&a)

	fmt.Print("Informe o valor do coeficiente B: ")
	fmt.Scan(&b)

	fmt.Print("Informe o valor do coeficiente C: ")
	fmt.Scan(&c)

	delta := (b * b) - (4 * a * c)

	fmt.Printf("O VALOR DE DELTA E = %.2f\n\n", delta)
}