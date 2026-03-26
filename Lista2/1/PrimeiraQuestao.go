package main
import ("fmt")

func main(){
	var numero int
	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&numero)
	
	if numero % 2 == 0 {
		fmt.Printf("O número %v é par.\n", numero)
	} else {
		fmt.Printf("O número %v é ímpar.\n", numero)
	}
}
