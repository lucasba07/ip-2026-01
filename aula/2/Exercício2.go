package main
import ("fmt")

func main(){
	var n1, n2, p1, p2 float64

	fmt.Print("Informe as duas notas: ")
	fmt.Scan(&n1, &n2)

	fmt.Print("Informe os dois pesos na ordem das notas: ")
	fmt.Scan(&p1, &p2)

	media := (n1 * p1) + (n2 * p2)

	fmt.Printf("A média ponderada das notas é: %v\n", media)
}