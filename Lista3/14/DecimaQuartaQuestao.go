package main
import "fmt"

func main(){
	var(
		n1 int
		n2 int
	)

	fmt.Print("Informe o primeiro número natural: ")
	fmt.Scanln(&n1)

	fmt.Print("Informe o segundo número natural: ")
	fmt.Scanln(&n2)

	if n1 == n2 || n1 < 0 || n2 < 0 || n1 > n2 {fmt.Println("VALORES INVÁLIDOS!")}

	for i:=(n1+1); i < n2; i++{
		if verificaPrimo(i){fmt.Printf("%v ", i)}
	}
}

func verificaPrimo(num int) bool{
	if num == 1 {return false}
	for x:=2; x < num; x++{
		if num % x == 0{
			return false
		} 
	}
	return true
}
