package main
import ("fmt")

func main(){
	var horas int

	fmt.Print("Informe a quantidade de horas em que você usou o charrete: ")
	fmt.Scan(&horas)
	fmt.Print("\n")

	switch horas % 3 {
		case 0:
			valorPagar := float32(10 * (horas / 3))
			fmt.Printf("O VALOR A PAGAR E = %.2f\n\n", valorPagar)
		case 1:
			valorPagar := float32(10 * (horas / 3)) + 5
			fmt.Printf("O VALOR A PAGAR E = %.2f\n\n", valorPagar)
		case 2:
			valorPagar := float32(10 * (horas / 3)) + 10
			fmt.Printf("O VALOR A PAGAR E = %.2f\n\n", valorPagar)
		}
}