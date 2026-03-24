package main
import ("fmt")

func main(){
	var l1, l2, l3 float64
	fmt.Print("Informe os lados do triângulo: ")
		fmt.Scan(&l1, &l2, &l3)

	switch{
		case l1 < (l2 + l3) && l2 < (l1 + l3) && l3 < (l1 + l2):
			if l1 == l2 && l2 == l3 {
				fmt.Println("O triângulo é Equilátero!")
			} else if (l1 == l2 && l2 != l3) || (l1 == l3 && l3 != l2) || (l2 == l3 && l3 != l1){
				fmt.Println("O triângulo é Isósceles!")
			} else {
				fmt.Println("O triângulo é Escaleno!")
			}
		default:
			fmt.Println("Valores inválidos, tente novamente!")
	}
}