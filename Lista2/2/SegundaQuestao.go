package main
import ("fmt")

func main(){
	var numero int
	
	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&numero)

	if numero > 0 {
		fmt.Printf("O número %v é positivo.\n", numero)
	} else if numero < 0 {
		fmt.Printf("O número %v é negativo.\n", numero)
	} else {
		fmt.Printf("O número é zero.\n")
	}
}
