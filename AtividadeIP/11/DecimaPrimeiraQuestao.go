package main
import ("fmt")

func main(){
	var num int

	fmt.Print("Informe um número inteiro qualquer: ")
	fmt.Scan(&num)
	fmt.Print("\n")

	if num % 5 == 0 && num % 3 == 0 {
		fmt.Println("O NUMERO E DIVISIVEL")
		fmt.Print("\n")
	} else {
		fmt.Println("O NUMERO NAO E DIVISIVEL")
		fmt.Print("\n")
	}
}